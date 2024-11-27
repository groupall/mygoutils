package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type BD struct {
	Dsn string
	Bd  string
}
type JWT struct {
	Secret string
}

type Config struct {
	Env                 string
	Service             string
	BD                  BD
	Jwt                 JWT
	DebugLevel          string
	RpcTimeout          int64
	HashCost            string
	LogUpdateServer     int64
	BrokerServer        string
	BrokerServerPort    string
	BrokerDSN           string
	BrokerBD            string
	BrokerGrpcTls       string
	BrokerGrpcCertFile  string
	BrokerGrpcCertKey   string
	BrokerGrpcClientKey string
	DirServer           string
	DirServerPort       string
	DirDSN              string
	DirBD               string
	DirGrpcTls          string
	DirGrpcCertFile     string
	DirGrpcCertKey      string
	DirGrpcClientKey    string
	LogServer           string
	LogServerPort       string
	LogDSN              string
	LogBD               string
	LogGrpcTls          string
	LogGrpcCertFile     string
	LogGrpcCertKey      string
	LogGrpcClientKey    string
	GeoServer           string
	GeoServerPort       string
	GeoDSN              string
	GeoBD               string
	GeoGrpcTls          string
	GeoGrpcCertFile     string
	GeoGrpcCertKey      string
	GeoGrpcClientKey    string
	AuthServer          string
	AuthServerPort      string
	AuthDSN             string
	AuthBD              string
	AuthGrpcTls         string
	AuthGrpcCertFile    string
	AuthGrpcCertKey     string
	AuthGrpcClientKey   string
	GrupoServer         string
	GrupoServerPort     string
	GrupoDSN            string
	GrupoBD             string
	GrupoGrpcTls        string
	GrupoGrpcCertFile   string
	GrupoGrpcCertKey    string
	GrupoGrpcClientKey  string
	DiccioServer        string
	DiccioServerPort    string
	DiccioDSN           string
	DiccioBD            string
	DiccioGrpcTls       string
	DiccioGrpcCertFile  string
	DiccioGrpcCertKey   string
	DiccioGrpcClientKey string
	MailServer          string
	MailServerPort      string
	MailDSN             string
	MailBD              string
	MailGrpcTls         string
	MailGrpcCertFile    string
	MailGrpcCertKey     string
	MailGrpcClientKey   string
	RedisServer         string
	RedisServerPort     string
	RedisPass           string
	RedisExpiration     string
	SendGridApi	string
}

type Svc struct {
	Servicio string
}

func (s *Svc) cargarEnvVariable(envvar string) string {
	valor := os.Getenv(envvar)
	if valor == "" {
		log.Printf("%s La variable %s está vacía.", s.Servicio, envvar)
	}
	return valor
}

func GetCFG(servicio string) (Config, error) {
	var cfg Config
	var err error
	var svc Svc
	svc.Servicio = servicio

	args := os.Args[1:]

	if len(args) > 0 {
		//Me dio un nombre de environment
		err4 := godotenv.Load(args[0])
		if err4 != nil {
			log.Printf("%s-Error loading env file: ", servicio, args[0])
		} else {
			log.Printf("%s-Cargó %s", servicio, args[0])
		}
	} else {
		// hago lo de default. Busco un .env
		err = godotenv.Load()
		if err != nil {
			// No había un .env.  Voy a buscar el local.env
			err2 := godotenv.Load("local.env")
			if err2 != nil {
				log.Printf("%s-Error loading local.env file", servicio)
			} else {
				log.Printf("%s-Cargó local.env", servicio)
			}
		} else {
			log.Printf("%s-Cargó .env", servicio)
		}
	}

	// s3Bucket := os.Getenv("S3_BUCKET")
	// secretKey := os.Getenv("SECRET_KEY")

	// cfg.Jwt.Secret = os.Getenv("JWT")

	logus, err := strconv.ParseInt(svc.cargarEnvVariable("LOG_UPDATE_SERVER"), 10, 64)
	if err != nil {
		log.Print("Error al convertir la variable LOG_UPDATE_SERVER")
		cfg.LogUpdateServer = 1 //son minutos
	} else {
		cfg.LogUpdateServer = logus
	}

	to, err := strconv.ParseInt(svc.cargarEnvVariable("RPC_TIMEOUT"), 10, 64)
	if err != nil {
		log.Print("Error al convertir la variable RPC_TIMEOUT")
		cfg.RpcTimeout = 2
	} else {
		cfg.RpcTimeout = to
	}

	
	cfg.SendGridApi = svc.cargarEnvVariable("SENDGRID_API_KEY")
	cfg.HashCost = svc.cargarEnvVariable("HASH_COST")
	cfg.DebugLevel = svc.cargarEnvVariable("DEBUG_LEVEL")

	cfg.Service = servicio

	cfg.Env = svc.cargarEnvVariable("ENVIRONMENT")

	// cfg.BD.Dsn  = cargarEnvVariable("DSN")

	// cfg.BD.Bd = cargarEnvVariable("BD")

	cfg.DirServer = svc.cargarEnvVariable("DIRGRPCSERVER")
	cfg.DirServerPort = svc.cargarEnvVariable("DIRGRPCSERVERPORT")
	cfg.DirDSN = svc.cargarEnvVariable("DIRDSN")
	cfg.DirBD = svc.cargarEnvVariable("DIRBD")
	cfg.DirGrpcTls = svc.cargarEnvVariable("DIRGRPCTLS")
	cfg.DirGrpcCertFile = svc.cargarEnvVariable("DIRGRPCTLSCERTFILE")
	cfg.DirGrpcCertKey = svc.cargarEnvVariable("DIRGRPCTLSCERTKEY")
	cfg.DirGrpcCertKey = svc.cargarEnvVariable("DIRGRPCTLSCERTKEY")
	cfg.DirGrpcClientKey = svc.cargarEnvVariable("DIRGRPCTLSCLIENTKEY")

	cfg.BrokerServer = svc.cargarEnvVariable("BROKERGRPCSERVER")
	cfg.BrokerServerPort = svc.cargarEnvVariable("BROKERGRPCSERVERPORT")
	cfg.BrokerDSN = svc.cargarEnvVariable("BROKERDSN")
	cfg.BrokerBD = svc.cargarEnvVariable("BROKERBD")
	cfg.BrokerGrpcTls = svc.cargarEnvVariable("BROKERGRPCTLS")
	cfg.BrokerGrpcCertFile = svc.cargarEnvVariable("BROKERGRPCTLSCERTFILE")
	cfg.BrokerGrpcCertKey = svc.cargarEnvVariable("BROKERGRPCTLSCERTKEY")
	cfg.BrokerGrpcCertKey = svc.cargarEnvVariable("BROKERGRPCTLSCERTKEY")
	cfg.BrokerGrpcClientKey = svc.cargarEnvVariable("BROKERGRPCTLSCLIENTKEY")

	cfg.LogServer = svc.cargarEnvVariable("LOGGRPCSERVER")
	cfg.LogServerPort = svc.cargarEnvVariable("LOGGRPCSERVERPORT")
	cfg.LogDSN = svc.cargarEnvVariable("LOGDSN")
	cfg.LogBD = svc.cargarEnvVariable("LOGBD")
	cfg.LogGrpcTls = svc.cargarEnvVariable("LOGGRPCTLS")
	cfg.LogGrpcCertFile = svc.cargarEnvVariable("LOGGRPCTLSCERTFILE")
	cfg.LogGrpcCertKey = svc.cargarEnvVariable("LOGGRPCTLSCERTKEY")
	cfg.LogGrpcCertKey = svc.cargarEnvVariable("LOGGRPCTLSCERTKEY")
	cfg.LogGrpcClientKey = svc.cargarEnvVariable("LOGGRPCTLSCLIENTKEY")

	cfg.AuthServer = svc.cargarEnvVariable("AUTHGRPCSERVER")
	cfg.AuthServerPort = svc.cargarEnvVariable("AUTHGRPCSERVERPORT")
	cfg.AuthDSN = svc.cargarEnvVariable("AUTHDSN")
	cfg.AuthBD = svc.cargarEnvVariable("AUTHBD")
	cfg.AuthGrpcTls = svc.cargarEnvVariable("AUTHGRPCTLS")
	cfg.AuthGrpcCertFile = svc.cargarEnvVariable("AUTHGRPCTLSCERTFILE")
	cfg.AuthGrpcCertKey = svc.cargarEnvVariable("AUTHGRPCTLSCERTKEY")
	cfg.AuthGrpcClientKey = svc.cargarEnvVariable("AUTHGRPCTLSCLIENTKEY")

	cfg.GrupoServer = svc.cargarEnvVariable("GRUPOGRPCSERVER")
	cfg.GrupoServerPort = svc.cargarEnvVariable("GRUPOGRPCSERVERPORT")
	cfg.GrupoDSN = svc.cargarEnvVariable("GRUPODSN")
	cfg.GrupoBD = svc.cargarEnvVariable("GRUPOBD")
	cfg.GrupoGrpcTls = svc.cargarEnvVariable("GRUPOGRPCTLS")
	cfg.GrupoGrpcCertFile = svc.cargarEnvVariable("GRUPOGRPCTLSCERTFILE")
	cfg.GrupoGrpcCertKey = svc.cargarEnvVariable("GRUPOGRPCTLSCERTKEY")
	cfg.GrupoGrpcClientKey = svc.cargarEnvVariable("GRUPOGRPCTLSCLIENTKEY")

	cfg.GeoServer = svc.cargarEnvVariable("GEOGRPCSERVER")
	cfg.GeoServerPort = svc.cargarEnvVariable("GEOGRPCSERVERPORT")
	cfg.GeoDSN = svc.cargarEnvVariable("GEODSN")
	cfg.GeoBD = svc.cargarEnvVariable("GEOBD")
	cfg.GeoGrpcTls = svc.cargarEnvVariable("GEOGRPCTLS")
	cfg.GeoGrpcCertFile = svc.cargarEnvVariable("GEOGRPCTLSCERTFILE")
	cfg.GeoGrpcCertKey = svc.cargarEnvVariable("GEOGRPCTLSCERTKEY")
	cfg.GeoGrpcClientKey = svc.cargarEnvVariable("GEOGRPCTLSCLIENTKEY")

	cfg.MailServer = svc.cargarEnvVariable("MAILGRPCSERVER")
	cfg.MailServerPort = svc.cargarEnvVariable("MAILGRPCSERVERPORT")
	cfg.MailDSN = svc.cargarEnvVariable("MAILDSN")
	cfg.MailBD = svc.cargarEnvVariable("MAILBD")
	cfg.MailGrpcTls = svc.cargarEnvVariable("MAILGRPCTLS")
	cfg.MailGrpcCertFile = svc.cargarEnvVariable("MAILGRPCTLSCERTFILE")
	cfg.MailGrpcCertKey = svc.cargarEnvVariable("MAILGRPCTLSCERTKEY")
	cfg.MailGrpcClientKey = svc.cargarEnvVariable("MAILGRPCTLSCLIENTKEY")

	cfg.DiccioServer = svc.cargarEnvVariable("DICCIOGRPCSERVER")
	cfg.DiccioServerPort = svc.cargarEnvVariable("DICCIOGRPCSERVERPORT")
	cfg.DiccioDSN = svc.cargarEnvVariable("DICCIODSN")
	cfg.DiccioBD = svc.cargarEnvVariable("DICCIOBD")
	cfg.DiccioGrpcTls = svc.cargarEnvVariable("DICCIOGRPCTLS")
	cfg.DiccioGrpcCertFile = svc.cargarEnvVariable("DICCIOGRPCTLSCERTFILE")
	cfg.DiccioGrpcCertKey = svc.cargarEnvVariable("DICCIOGRPCTLSCERTKEY")
	cfg.DiccioGrpcClientKey = svc.cargarEnvVariable("DICCIOGRPCTLSCLIENTKEY")

	cfg.RedisServer = svc.cargarEnvVariable("REDISSERVER")
	cfg.RedisServerPort = svc.cargarEnvVariable("REDISSERVERPORT")
	cfg.RedisExpiration = svc.cargarEnvVariable("REDISEXPIRATION")

	return cfg, nil

}
