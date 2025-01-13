package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee struct {
	FirstName string
	LastName  string
	Age       int
	Salary    float64
}

func TestSortEmployees(t *testing.T) {
	testCases := []struct {
		name      string
		employees []Employee
		expected  []Employee
	}{
		{
			name: "Sort by LastName, FirstName, Salary Desc",
			employees: []Employee{
				{FirstName: "Alice", LastName: "Smith", Age: 30, Salary: 70000},
				{FirstName: "Bob", LastName: "Johnson", Age: 35, Salary: 80000},
				{FirstName: "Charlie", LastName: "Smith", Age: 28, Salary: 70000},
				{FirstName: "Alice", LastName: "Johnson", Age: 32, Salary: 90000},
			},
			expected: []Employee{
				{FirstName: "Alice", LastName: "Johnson", Age: 32, Salary: 90000},
				{FirstName: "Bob", LastName: "Johnson", Age: 35, Salary: 80000},
				{FirstName: "Alice", LastName: "Smith", Age: 30, Salary: 70000},
				{FirstName: "Charlie", LastName: "Smith", Age: 28, Salary: 70000},
			},
		},
		{
			name: "Sort with identical LastNames",
			employees: []Employee{
				{FirstName: "Alice", LastName: "Smith", Age: 30, Salary: 70000},
				{FirstName: "Bob", LastName: "Smith", Age: 35, Salary: 80000},
				{FirstName: "Charlie", LastName: "Smith", Age: 28, Salary: 70000},
			},
			expected: []Employee{
				{FirstName: "Bob", LastName: "Smith", Age: 35, Salary: 80000},
				{FirstName: "Alice", LastName: "Smith", Age: 30, Salary: 70000},
				{FirstName: "Charlie", LastName: "Smith", Age: 28, Salary: 70000},
			},
		},
		{
			name: "Sort with identical Salaries",
			employees: []Employee{
				{FirstName: "Alice", LastName: "Smith", Age: 30, Salary: 70000},
				{FirstName: "Bob", LastName: "Smith", Age: 35, Salary: 70000},
				{FirstName: "Charlie", LastName: "Smith", Age: 28, Salary: 70000},
			},
			expected: []Employee{
				{FirstName: "Alice", LastName: "Smith", Age: 30, Salary: 70000},
				{FirstName: "Bob", LastName: "Smith", Age: 35, Salary: 70000},
				{FirstName: "Charlie", LastName: "Smith", Age: 28, Salary: 70000},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var employees = tc.employees
			got := make([]Employee, len(tc.employees))
			copy(got, tc.employees)

			sort.Slice(got, func(i, j int) bool {
				return got[i].LastName < got[j].LastName ||
					got[i].FirstName < got[j].FirstName ||
					got[i].Salary > got[j].Salary
			})

			for index := range employees {
				assert.Equal(t, tc.expected[index].FirstName, got[index].FirstName, "FirstName is incorrect")
				assert.Equal(t, tc.expected[index].LastName, got[index].LastName, "LastName is incorrect")
				assert.Equal(t, tc.expected[index].Age, got[index].Age, "Age is incorrect")
				assert.Equal(t, tc.expected[index].Salary, got[index].Salary, "Salary is incorrect")

			}
		})

	}
}
