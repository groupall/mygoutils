package utils

import (
	"context"
	"fmt"
	"log"
	"mirrioba/dir"
	"os/exec"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type ServerDir struct {
	Server string
	Port   string
}

func GrpcDial(servicio, url, usaTLS, clientKey string, block bool) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{}

	if usaTLS == "S" {
		creds, err := credentials.NewClientTLSFromFile(clientKey, "")
		if err != nil {
			log.Printf("%s-Error en la carga de certificados creación del cliente TLS, %v", servicio, err)
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		// uso insecure
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	// Le agrego que se bloquee hasta que se conecte....
	if block {
		opts = append(opts, grpc.WithBlock())
	}

	return grpc.Dial(url, opts...)

}

func GrpcCreateServer(servicio string, usaTLS string, certFile string, certKey string) (*grpc.Server, error) {

	opts := []grpc.ServerOption{}

	if usaTLS == "S" {
		log.Printf("%s iniciando servidor seguro..", servicio)
		creds, err := credentials.NewServerTLSFromFile(certFile, certKey)
		if err != nil {
			log.Printf("%s-Error en la carga de certificados creación del server TLS, %v", servicio, err)
			return nil, err
		}
		opts = append(opts, grpc.Creds(creds))
	}
	return grpc.NewServer(opts...), nil // los ... convierten el array en un conjunto de parametros.
}

func LevantarServidor(servicio string) {
	// log.Printf("Levantando servicio de: &s", servicio)
	var comando string

	switch servicio {
	case "DirSvc":
		comando = "./dirApp.o"
	case "LogsSvc":
		comando = "./logApp.o"
	case "MailSvc":
		comando = "./mailApp.o"
	case "GruposSvc":
		comando = "./grupoApp.o"
	case "GeoSvc":
		comando = "./geoApp.o"
	case "BrokerSvc":
		comando = "./brokerApp.o"
	case "AuthSvc":
		comando = "./authApp.o"
	default:
		log.Printf("Servicio a levantar no reconocido: %s", servicio)
		return
	}
	cmd := exec.Command(comando)
	if err := cmd.Run(); err != nil {
		log.Printf("Error al ejecutar el comando:%s, %v ", comando, err)
	}
}

func GetServerDir(servicio, serverDestino, portDestino, tls, clientKey string, RpcTimeout int64,
	clientConnect func(grpc.ClientConnInterface) dir.DirServiceClient) (ServerDir, error) {

	var s ServerDir

	conn, err := GrpcDial(servicio, fmt.Sprintf("%s:%s", serverDestino, portDestino), tls, clientKey, true)

	if err != nil {
		// voy a ver si el error es del tipo GRPC
		e, ok := status.FromError(err)
		if ok {
			log.Printf("%sNo se pudo conectar al grpc dir:%s, Cod:%s ", servicio, e.Message(), e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Print("Invalid arguments")
			}
		} else {
			//Non grpc error
			log.Printf("%sNo se pudo conectar al grpc log:%s ", servicio, err.Error())
		}
		return s, err
	}

	defer conn.Close()

	c := clientConnect(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(RpcTimeout))

	defer cancel()

	log.Printf("voy a mandar el pedido de ServiceData")
	v, err := c.ServiceData(ctx, &dir.Srv{Servicio: servicio})
	if err != nil {
		log.Printf("Error al mandar\n %v \n", err)
		return s, err
	}
	s.Server = v.Host
	s.Port = v.Port
	return s, nil

}





