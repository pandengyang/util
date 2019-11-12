package CollectionJSON

import (
	"testing"
)

func BenchmarkCollection(b *testing.B) {
	task1 := Task{1, "买书", "2018-09-23", false}
	task2 := Task{2, "吃月饼", "2018-10-23", false}
	tasks := [] interface{} {
		task1, task2,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Items(tasks, pTaskTemplateStr)
	}
}

func BenchmarkItem(b *testing.B) {
	task1 := Task{1, "买书", "2018-09-23", false}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Item(task1, pTaskTemplateStr)
	}
}

func BenchmarkQueries(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Queries(pTaskTemplateStr)
	}
}

func BenchmarkTemplate(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Template(pTaskTemplateStr)
	}
}

// go test -v -run="none" -bench="BenchmarkSprintf" -benchtime="3s"
// go test -v -run="none" -bench=. -benchtime="3s" -benchmem
