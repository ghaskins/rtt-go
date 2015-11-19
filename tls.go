package main

import (
	"crypto/tls"
	"os"
)

var certPEM = `-----BEGIN CERTIFICATE-----
MIID0zCCArugAwIBAgIJAOCBLxSQGp7NMA0GCSqGSIb3DQEBBQUAME8xCzAJBgNV
BAYTAlVTMRYwFAYDVQQIEw1NYXNzYWNodXNldHRzMQ8wDQYDVQQHEwZIb2xkZW4x
FzAVBgNVBAMTDmxvY2FsaG9zdDoyMDAxMB4XDTE1MDkwOTIwMTk0MloXDTI1MDkw
NjIwMTk0MlowTzELMAkGA1UEBhMCVVMxFjAUBgNVBAgTDU1hc3NhY2h1c2V0dHMx
DzANBgNVBAcTBkhvbGRlbjEXMBUGA1UEAxMObG9jYWxob3N0OjIwMDEwggEiMA0G
CSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC21kXPF/msRNbZdZMr4lo4yo/kW3dX
x4VWStFxOdDNCwv3L4CZAgKF5qAiFQCjmHUdWAfkV4S3zWIZdRKWN5Dpu3GcRntR
NSmifANOyE6Egx06IGu08jFqGPf0n8LANc7UCiQXyroF+Vfw3Q/cVVgAplHYXKYF
YxMUzj9rSxj2U4jxB4YkTPNAiTvKWcOrZLHfchoXdU9l5b6i0dD2LREaltjzUXbC
c4U13mGmrilijC8F8tahQfKGPcxpf+22eIcFSzXkO8Xz+oPgR2jfffK/P29F8nAG
GdxJpaKIlP8HJoSZS+2D6GbdRfvtuah6MKe17ZOLizrNPnWW2XhptRvJAgMBAAGj
gbEwga4wHQYDVR0OBBYEFJQES7YrqkRt68gpsu2bicylmXnfMH8GA1UdIwR4MHaA
FJQES7YrqkRt68gpsu2bicylmXnfoVOkUTBPMQswCQYDVQQGEwJVUzEWMBQGA1UE
CBMNTWFzc2FjaHVzZXR0czEPMA0GA1UEBxMGSG9sZGVuMRcwFQYDVQQDEw5sb2Nh
bGhvc3Q6MjAwMYIJAOCBLxSQGp7NMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEF
BQADggEBAKYcVS+D/YBr0KY+yRqfeQtUjBYDnoEjkGga6oaTe8GjfJuyG6cIKZn1
zCgHRg9EFvDE0x9i2G643WFPtN22MhkRU2ES4hpRQ89A3SPbMhfDKB4wK5MqTGVM
Q1K6J+lvguc512Vu5QdhRzfYpMORGEK9dpH2F7za5lwGHCIbgT8NR3ipcWDO7hnk
93zFNoX0fNsXmpobNp5E34AAvWg8okP9lnIHUCoCqhFHtzKGjYB1QJuhzDA7DpK5
yhkh+rdQiIunXjcPSJparq+HsKet4PpiuRqSjvBQpzpKefocFIbJCtqVQr5KQ92+
g0KDfOgDRDbtPIFg0QQVPABqKgfeg/o=
-----END CERTIFICATE-----
`

var keyPEM = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAttZFzxf5rETW2XWTK+JaOMqP5Ft3V8eFVkrRcTnQzQsL9y+A
mQICheagIhUAo5h1HVgH5FeEt81iGXUSljeQ6btxnEZ7UTUponwDTshOhIMdOiBr
tPIxahj39J/CwDXO1AokF8q6BflX8N0P3FVYAKZR2FymBWMTFM4/a0sY9lOI8QeG
JEzzQIk7ylnDq2Sx33IaF3VPZeW+otHQ9i0RGpbY81F2wnOFNd5hpq4pYowvBfLW
oUHyhj3MaX/ttniHBUs15DvF8/qD4Edo333yvz9vRfJwBhncSaWiiJT/ByaEmUvt
g+hm3UX77bmoejCnte2Ti4s6zT51ltl4abUbyQIDAQABAoIBAAhcciLPUN6c40pc
gxdtqXyRXtUtjZ6ZBDL3Bu7rmu3VhH2QUYwtuFnROx9z4GyayzhFT08U5X/WR5cs
cGYoMltN9BsWagtDqBzDGQ+gP1718/81fldu/+1/KM+qAqjxjPzLXe0smycsyOW9
5jzINlcJRBLl4b27UyMPb30DmSeihGjr5+BI4XJqJLAgAYg8DM9rts13JJX3YQu9
IcitLk7FK5VCkeCk3O7p+lcIz/PiobmdIypmolPlubSfVEs2g7MHzGsoP+JHndLf
O8phGeT7UMVvU3ssXguPEQV6FMzpEtsufp8MG9VYZd+yY2IcUAOvVZmVKseVizWH
1+t6gZUCgYEA8NeSp3mOkVXDN7VKcGEYkMkfPbLiIoADAdehD7Ex6ApQejT8eumQ
f613KFIHva+kK5QZYx0/uv+cCWRBK/F2w471Tp/l0nneb5kFSvYjxQa1FzOgHpgT
nNMZxVYF27VMwsSaletSm4Gj+ai9YUdc/H6fdSYbUQMFmwpmXjP7CnsCgYEAwlgg
fkxgcd17sV+981grwHA5jtE7cinL7fwexbpDmMukjoqOHf8XVYD6uG7GdC5fNLXa
wFi37P7PwEHjtSqr7NgDGCFXp+vuHXY70x4mehQBhjgRfg+mTdHZdctdUK2EBEPK
5jL1utVVthFewlRsIW+Q+i+Ozcs4MCHB0esI0YsCgYBc3HGyWz8qMGwt7Zu/CuEC
6lk+W9uvO7ZtHmv+de7tLhTrmcSDy8yoPgUEqeRMMg3Vs6u6OIvbGTVbtakfPWHy
cwuIfkSJy+2FD/YnehI+pKBsSr6BLVfajtaP7OQjW5s2OcH07iASz4CfAX3LpU1o
GZZ3//JmYQjnR7JCvj4cQwKBgBzRYJAZ1ztLMmpM3ifVbTdt7LsGv5/gHsM9XXrI
cfmYSOByHzzHdLhTFdp/YqIbtjZkuPlIro1QA5JostFMcI4AQgUz/IGD9J0dySU3
FVGI4ej7K2zb+TcU5QAYlc++yxKu0slryRBZTgpfbQO37QjQOFlA8BYICl7oworU
JUydAoGBAIRC/T8A1TGmvpf2ibfJbqxf2peqrLJDj5d4uMwg3k47HAKI/93cw2Y1
aDW5d4LZH1Pxkj2QKMcT+mQGV8GOAcRHCF+IH3Og85OORRl2lmNW+d5qXhyAT2ni
yGiXlnQv0rl5bbW2XeH4aU4bwOpOF+V3Svo62IxrLYyYm3ldn+6Z
-----END RSA PRIVATE KEY-----
`

func newServerConfig(self *tls.Certificate) *tls.Config {
	config := &tls.Config{
		Certificates:       make([]tls.Certificate, 1),
	}

	config.Certificates[0] = *self

	return config
}

func NewTLS(payloadLen int) func() {

	laddr := "./rtt-go-tls.unix"
	cert,err := tls.X509KeyPair([]byte(certPEM), []byte(keyPEM))

	// Start our listener in common-context so we don't race with the registration
	listener, err := tls.Listen("unix", laddr, newServerConfig(&cert))
	if err != nil {
		panic(err)
	}

	go func () {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		buf := make([]byte, payloadLen)

		for {
			conn.Read(buf)
			conn.Write(buf)
		}
	}()

	conn, err := tls.Dial("unix", laddr, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		panic(err)
	}

	os.Remove(laddr)

	buf := make([]byte, payloadLen)

	return func() {
			conn.Write(buf)
			conn.Read(buf)
	}
}
