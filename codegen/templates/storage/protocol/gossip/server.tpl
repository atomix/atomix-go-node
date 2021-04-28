{{- $serviceType := printf "%sServiceType" .Generator.Prefix }}
{{- $server := printf "%sServer" .Generator.Prefix }}
{{- $service := printf "%sService" .Generator.Prefix }}
package {{ .Package.Name }}

import (
	"context"
	"github.com/atomix/go-framework/pkg/atomix/errors"
	"github.com/atomix/go-framework/pkg/atomix/logging"
	"github.com/atomix/go-framework/pkg/atomix/storage/protocol/gossip"
	"google.golang.org/grpc"
	{{- $package := .Package }}
	{{- range .Imports }}
	{{ .Alias }} {{ .Path | quote }}
	{{- end }}
	{{- range .Primitive.Methods }}
	{{- if .Request.IsStream }}
	{{ import "io" }}
	{{- end }}
	{{- end }}
)

// Register{{ $server }} registers the primitive on the given node
func Register{{ $server }}(node *gossip.Node) {
	node.RegisterServer(func(server *grpc.Server, manager *gossip.Manager) {
		{{ .Primitive.Type.Package.Alias }}.Register{{ .Primitive.Type.Name }}Server(server, new{{ $server }}(manager))
	})
}

func new{{ $server }}(manager *gossip.Manager) {{ .Primitive.Type.Package.Alias }}.{{ .Primitive.Type.Name }}Server {
	return &{{ $server }}{
		manager: manager,
		log: logging.GetLogger("atomix", "protocol", "gossip", {{ .Primitive.Name | lower | quote }}),
	}
}

{{- $primitive := .Primitive }}
{{- $serviceInt := printf "%sService" .Generator.Prefix }}
type {{ $server }} struct {
    manager *gossip.Manager
	log logging.Logger
}

{{- define "type" -}}
{{- if .Package.Import -}}
{{- printf "%s.%s" .Package.Alias .Name -}}
{{- else -}}
{{- .Name -}}
{{- end -}}
{{- end -}}

{{- define "cast" }}.(*{{ template "type" .Field.Type }}){{ end }}

{{- define "field" }}
{{- $path := .Field.Path }}
{{- range $index, $element := $path -}}
{{- if eq $index 0 -}}
{{- if isLast $path $index -}}
{{- if $element.Type.IsPointer -}}
.Get{{ $element.Name }}()
{{- else -}}
.{{ $element.Name }}
{{- end -}}
{{- else -}}
{{- if $element.Type.IsPointer -}}
.Get{{ $element.Name }}().
{{- else -}}
.{{ $element.Name }}.
{{- end -}}
{{- end -}}
{{- else -}}
{{- if isLast $path $index -}}
{{- if $element.Type.IsPointer -}}
    Get{{ $element.Name }}()
{{- else -}}
    {{ $element.Name -}}
{{- end -}}
{{- else -}}
{{- if $element.Type.IsPointer -}}
    Get{{ $element.Name }}().
{{- else -}}
    {{ $element.Name }}.
{{- end -}}
{{- end -}}
{{- end -}}
{{- end -}}
{{- end }}

{{- define "ref" -}}
{{- if not .Field.Type.IsPointer }}&{{ end }}
{{- end }}

{{- define "val" -}}
{{- if .Field.Type.IsPointer }}*{{ end }}
{{- end }}

{{- define "optype" }}
{{- if .Type.IsCommand -}}
Command
{{- else if .Type.IsQuery -}}
Query
{{- end -}}
{{- end }}

{{- range .Primitive.Methods }}
{{- $method := . }}
{{ if and .Request.IsDiscrete .Response.IsDiscrete }}
func (s *{{ $server }}) {{ .Name }}(ctx context.Context, request *{{ template "type" .Request.Type }}) (*{{ template "type" .Response.Type }}, error) {
	s.log.Debugf("Received {{ .Request.Type.Name }} %+v", request)
	s.manager.AddRequestHeaders({{ template "ref" .Request.Headers }}request{{ template "field" .Request.Headers }})
	partition, err := s.manager.Partition(gossip.PartitionID(request{{ template "field" .Request.Headers }}.PartitionID))
    if err != nil {
        s.log.Errorf("Request {{ .Request.Type.Name }} %+v failed: %v", request, err)
        return nil, err
    }

    serviceID := gossip.ServiceId{
        Type:      gossip.ServiceType(request{{ template "field" .Request.Headers }}.PrimitiveID.Type),
        Namespace: request{{ template "field" .Request.Headers }}.PrimitiveID.Namespace,
        Name:      request{{ template "field" .Request.Headers }}.PrimitiveID.Name,
    }

    service, err := partition.GetService(ctx, serviceID)
    if err != nil {
        s.log.Errorf("Request {{ .Request.Type.Name }} %+v failed: %v", request, err)
        return nil, errors.Proto(err)
    }

    response, err := service.({{ $service }}).{{ .Name }}(ctx, request)
    if err != nil {
        s.log.Errorf("Request {{ .Request.Type.Name }} %+v failed: %v", request, err)
        return nil, errors.Proto(err)
    }
    s.manager.AddResponseHeaders({{ template "ref" .Response.Headers }}response{{ template "field" .Response.Headers }})
	s.log.Debugf("Sending {{ .Response.Type.Name }} %+v", response)
	return response, nil
}
{{ else if .Response.IsStream }}
func (s *{{ $server }}) {{ .Name }}(request *{{ template "type" .Request.Type }}, srv {{ template "type" $primitive.Type }}_{{ .Name }}Server) error {
    s.log.Debugf("Received {{ .Request.Type.Name }} %+v", request)
	s.manager.AddRequestHeaders({{ template "ref" .Request.Headers }}request{{ template "field" .Request.Headers }})

	partition, err := s.manager.Partition(gossip.PartitionID(request{{ template "field" .Request.Headers }}.PartitionID))
    if err != nil {
        s.log.Errorf("Request {{ .Request.Type.Name }} %+v failed: %v", request, err)
        return errors.Proto(err)
    }

    serviceID := gossip.ServiceId{
        Type:      gossip.ServiceType(request{{ template "field" .Request.Headers }}.PrimitiveID.Type),
        Namespace: request{{ template "field" .Request.Headers }}.PrimitiveID.Namespace,
        Name:      request{{ template "field" .Request.Headers }}.PrimitiveID.Name,
    }

    service, err := partition.GetService(srv.Context(), serviceID)
    if err != nil {
        s.log.Errorf("Request {{ .Request.Type.Name }} %+v failed: %v", request, err)
        return err
    }

    responseCh := make(chan {{ template "type" .Response.Type }})
    errCh := make(chan error)
    go func() {
        err := service.({{ $service }}).{{ .Name }}(srv.Context(), request, responseCh)
        if err != nil {
            errCh <- err
        }
        close(errCh)
    }()

    for {
        select {
        case response, ok := <-responseCh:
            if ok {
                s.manager.AddResponseHeaders({{ template "ref" .Response.Headers }}response{{ template "field" .Response.Headers }})
                s.log.Debugf("Sending {{ .Response.Type.Name }} %v", response)
                err = srv.Send(&response)
                if err != nil {
                    s.log.Errorf("Request {{ .Request.Type.Name }} %+v failed: %v", request, err)
                    return errors.Proto(err)
                }
            } else {
                s.log.Debugf("Finished {{ .Request.Type.Name }} %+v", request)
                return nil
            }
        case <-srv.Context().Done():
            s.log.Debugf("Finished {{ .Request.Type.Name }} %+v", request)
            return nil
        }
    }
}
{{ end }}
{{- end }}
