package utils

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetDate() time.Time {

	return time.Now().UTC()

}

func GetNewUUID() string {
	return uuid.NewV4().String()
}

func GetCurrentFuncName(skip int) string {
	pc, _, _, _ := runtime.Caller(1 + skip)
	// return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
	return runtime.FuncForPC(pc).Name()
}

func GetNextAvailablePort(servicio string, initial string) string {

	port, err := strconv.ParseUint(initial, 10, 16)
	if err != nil {
		fmt.Printf("%s-Invalid port %s: %v - %v", servicio, initial, err, os.Stderr)
		os.Exit(1)
	}

	retries := 1000
	log.Printf("voy a probar con el puerto %d", port)
	log.Print("dir:", fmt.Sprintf(":%s", fmt.Sprint(port)))
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", fmt.Sprint(port)))
	for err != nil && retries > 0 {
		log.Printf("Error en %d ... buscando nuevo puerto", port)
		retries--
		port++
		ln, err = net.Listen("tcp", fmt.Sprintf(":%s", fmt.Sprint(port)))
	}

	err = ln.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't stop listening on port %q: %s", port, err)
		os.Exit(1)
	}
	return fmt.Sprint(port)

}

func GetSiguienteEntero(valorActual, comienzo, total int64) int64 {
	proximo := valorActual + 1
	if proximo >= total {
		return comienzo
	}
	return proximo
}

func GetMyPublicIP() net.IP {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP
			}
		}
	}

	return nil
}

func GetMyInternalIP() string {
	return "localhost"
	conn, err := net.Dial("udp", "8.8.8.8:80")
	MiError("net.Dial: ", err, false)
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

func GetIpFromHost(host string) []net.IP {
	add, err := net.LookupIP(host)

	if err != nil {

		fmt.Println("host is unknown")
		return nil

	} else {

		fmt.Println("This is the IP address: ", add)
		return nil

	}
	return add
}

func GetServerUtilization() string {
	cpu_usage := 20.0
	mem_usage := 70.0
	formula := 0.7*cpu_usage + 0.3*mem_usage
	return fmt.Sprint(formula)
}

// func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
// 	wrapper := make(map[string]interface{})

// 	wrapper[wrap] = data

// 	js, err := json.Marshal(wrapper)
// 	if err != nil {
// 		return err
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	w.Write(js)

// 	return nil
// }

// func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) {
// 	statusCode := http.StatusBadRequest

// 	if len(status) > 0 {
// 		statusCode = status[0]
// 	}

// 	type jsonError struct {
// 		Message string `json:"message"`
// 	}

// 	theError := jsonError{
// 		Message: err.Error(),
// 	}

// 	fmt.Println(theError)
// 	app.writeJSON(w, statusCode, theError, "error")

// }
