package kestest_test

import (
	"testing"

	"github.com/minio/kes/internal/keystore/mem"
	"github.com/minio/kes/kv"
)

func TestGatewayMem(t *testing.T) {
	ctx, cancel := testingContext(t)
	defer cancel()

	store := kv.Store[string, []byte](&mem.Store{})

	t.Run("Metrics", func(t *testing.T) { testMetrics(ctx, store, t) })
	t.Run("APIs", func(t *testing.T) { testAPIs(ctx, store, t) })
	t.Run("CreateKey", func(t *testing.T) { testCreateKey(ctx, store, t) })
	t.Run("ImportKey", func(t *testing.T) { testImportKey(ctx, store, t) })
	t.Run("GenerateKey", func(t *testing.T) { testGenerateKey(ctx, store, t) })
	t.Run("EncryptKey", func(t *testing.T) { testEncryptKey(ctx, store, t) })
	t.Run("DecryptKey", func(t *testing.T) { testDecryptKey(ctx, store, t) })
	t.Run("DecryptKeyAll", func(t *testing.T) { testDecryptKeyAll(ctx, store, t) })
	t.Run("DescribePolicy", func(t *testing.T) { testDescribePolicy(ctx, store, t) })
	t.Run("GetPolicy", func(t *testing.T) { testGetPolicy(ctx, store, t) })
	t.Run("SelfDescribe", func(t *testing.T) { testSelfDescribe(ctx, store, t) })
}
