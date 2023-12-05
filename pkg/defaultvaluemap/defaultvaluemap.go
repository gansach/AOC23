// pkg/mydefaultmap/defaultvaluemap.go

package defaultvaluemap

type KeyValuePair[K comparable, V any] struct {
	Key   K
	Value V
}

type DefaultValueMap[K comparable, V any] struct {
	data         map[K]V
	defaultValue V
}

func New[K comparable, V any](defaultValue V) *DefaultValueMap[K, V] {
	return &DefaultValueMap[K, V]{
		data:         make(map[K]V),
		defaultValue: defaultValue,
	}
}

func (m *DefaultValueMap[K, V]) Get(key K) V {
	if value, ok := m.data[key]; ok {
		return value
	}
	return m.defaultValue
}

func (m *DefaultValueMap[K, V]) Set(key K, value V) {
	m.data[key] = value
}

func (m *DefaultValueMap[K, V]) Range() <-chan KeyValuePair[K, V] {
	resultCh := make(chan KeyValuePair[K, V])

	go func() {
		defer close(resultCh)
		for key, value := range m.data {
			resultCh <- KeyValuePair[K, V]{Key: key, Value: value}
		}
	}()

	return resultCh
}
