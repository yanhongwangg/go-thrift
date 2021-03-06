// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"github.com/apache/thrift/lib/go/thrift"
	"sync"
)

var _ = sync.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  SyncStatus check(ConfirmInfo info)")
  fmt.Fprintln(os.Stderr, "  SyncStatus startSync()")
  fmt.Fprintln(os.Stderr, "  SyncStatus init(string storageGroupName)")
  fmt.Fprintln(os.Stderr, "  SyncStatus syncDeletedFileName(string fileName)")
  fmt.Fprintln(os.Stderr, "  SyncStatus initSyncData(string filename)")
  fmt.Fprintln(os.Stderr, "  SyncStatus syncData(string buff)")
  fmt.Fprintln(os.Stderr, "  SyncStatus checkDataMD5(string md5)")
  fmt.Fprintln(os.Stderr, "  SyncStatus endSync()")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := sync.NewSyncServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "check":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Check requires 1 args")
      flag.Usage()
    }
    arg18 := flag.Arg(1)
    mbTrans19 := thrift.NewTMemoryBufferLen(len(arg18))
    defer mbTrans19.Close()
    _, err20 := mbTrans19.WriteString(arg18)
    if err20 != nil {
      Usage()
      return
    }
    factory21 := thrift.NewTJSONProtocolFactory()
    jsProt22 := factory21.GetProtocol(mbTrans19)
    argvalue0 := sync.NewConfirmInfo()
    err23 := argvalue0.Read(jsProt22)
    if err23 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Check(context.Background(), value0))
    fmt.Print("\n")
    break
  case "startSync":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "StartSync requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.StartSync(context.Background()))
    fmt.Print("\n")
    break
  case "init":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Init requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.Init(context.Background(), value0))
    fmt.Print("\n")
    break
  case "syncDeletedFileName":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SyncDeletedFileName requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.SyncDeletedFileName(context.Background(), value0))
    fmt.Print("\n")
    break
  case "initSyncData":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "InitSyncData requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.InitSyncData(context.Background(), value0))
    fmt.Print("\n")
    break
  case "syncData":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SyncData requires 1 args")
      flag.Usage()
    }
    argvalue0 := []byte(flag.Arg(1))
    value0 := argvalue0
    fmt.Print(client.SyncData(context.Background(), value0))
    fmt.Print("\n")
    break
  case "checkDataMD5":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CheckDataMD5 requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.CheckDataMD5(context.Background(), value0))
    fmt.Print("\n")
    break
  case "endSync":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "EndSync requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.EndSync(context.Background()))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
