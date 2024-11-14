package pki

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientCertNoSpiffeIdFromIntermediate(t *testing.T) {
	out, errOut := streams()
	svr := NewCmdPKICreateClient(out, errOut)
	name := uuid.New().String()
	args := []string{
		fmt.Sprintf("--pki-root=%s", where),
		fmt.Sprintf("--ca-name=%s", intCaNameWithoutSpiffeIdName),
		fmt.Sprintf("--client-name=%s", name),
		fmt.Sprintf("--client-file=%s", name),
		fmt.Sprintf("--dns=%s", "localhost,dns.entry"),
		fmt.Sprintf("--ip=%s", "127.0.0.1,::1"),
	}

	svr.SetArgs(args)
	svrErr := svr.Execute()
	if svrErr != nil {
		logrus.Fatal(svrErr)
	}

	bundle, e := testPki.GetBundle(intCaNameWithoutSpiffeIdName, name)
	assert.NotNil(t, bundle)
	assert.Nil(t, e)

	assert.Contains(t, bundle.Cert.DNSNames, "dns.entry")
	assert.Contains(t, bundle.Cert.DNSNames, "localhost")
	ips := ipsAsStrings(bundle.Cert.IPAddresses)
	assert.Contains(t, ips, "127.0.0.1")
	assert.Contains(t, ips, "::1")
	assert.Nil(t, bundle.Cert.URIs)
}

func TestClientCertSpiffeIdFromIntermediate(t *testing.T) {
	out, errOut := streams()
	svr := NewCmdPKICreateClient(out, errOut)
	name := uuid.New().String()
	args := []string{
		fmt.Sprintf("--pki-root=%s", where),
		fmt.Sprintf("--ca-name=%s", intCaNameWithSpiffeIdName),
		fmt.Sprintf("--client-name=%s", name),
		fmt.Sprintf("--client-file=%s", name),
		fmt.Sprintf("--dns=%s", "localhost,dns.entry"),
		fmt.Sprintf("--ip=%s", "127.0.0.1,::1"),
	}

	svr.SetArgs(addSpiffeArg("/some/path", args))
	svrErr := svr.Execute()
	if svrErr != nil {
		logrus.Fatal(svrErr)
	}

	bundle, e := testPki.GetBundle(intCaNameWithSpiffeIdName, name)
	assert.NotNil(t, bundle)
	assert.Nil(t, e)
	urls := URLSlice(bundle.Cert.URIs)

	assert.Contains(t, bundle.Cert.DNSNames, "dns.entry")
	assert.Contains(t, bundle.Cert.DNSNames, "localhost")
	ips := ipsAsStrings(bundle.Cert.IPAddresses)
	assert.Contains(t, ips, "127.0.0.1")
	assert.Contains(t, ips, "::1")
	assert.Contains(t, urls.Hosts(), rootCaWithSpiffeIdName)
	assert.Contains(t, urls.Paths(), "/some/path")
}

func TestClientCertNoSpiffeIdFromIntermediateAddSpiffeId(t *testing.T) {
	out, errOut := streams()
	svr := NewCmdPKICreateClient(out, errOut)
	name := uuid.New().String()
	args := []string{
		fmt.Sprintf("--pki-root=%s", where),
		fmt.Sprintf("--ca-name=%s", intCaNameWithoutSpiffeIdName),
		fmt.Sprintf("--client-name=%s", name),
		fmt.Sprintf("--client-file=%s", name),
	}

	sid := "spiffe://not-from-ca/the-path"
	svr.SetArgs(addSpiffeArg(sid, args))
	svrErr := svr.Execute()
	if svrErr != nil {
		logrus.Fatal(svrErr)
	}

	bundle, e := testPki.GetBundle(intCaNameWithoutSpiffeIdName, name)
	assert.NotNil(t, bundle)
	assert.Nil(t, e)

	assert.Contains(t, urisAsStrings(bundle.Cert.URIs), sid)
}

func TestClientCertSpiffeIdFromIntermediateAddSpiffeId(t *testing.T) {
	out, errOut := streams()
	svr := NewCmdPKICreateClient(out, errOut)
	name := uuid.New().String()
	args := []string{
		fmt.Sprintf("--pki-root=%s", where),
		fmt.Sprintf("--ca-name=%s", intCaNameWithSpiffeIdName),
		fmt.Sprintf("--client-name=%s", name),
		fmt.Sprintf("--client-file=%s", name),
	}

	sid := "spiffe://from-ca/the-path"
	svr.SetArgs(addSpiffeArg(sid, args))
	svrErr := svr.Execute()
	if svrErr != nil {
		logrus.Fatal(svrErr)
	}

	bundle, e := testPki.GetBundle(intCaNameWithSpiffeIdName, name)
	assert.NotNil(t, bundle)
	assert.Nil(t, e)

	assert.Contains(t, urisAsStrings(bundle.Cert.URIs), sid)
}
