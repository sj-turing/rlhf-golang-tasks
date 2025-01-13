package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee struct {
	FirstName string
	LastName  string
	Age       int
	Salary    int
}

func (e Employee) String() string {
	return fmt.Sprintf("%s %s - Age: %d, Salary: %d", e.FirstName, e.LastName, e.Age, e.Salary)
}

func TestSortEmployees(t *testing.T) {
	testCases := []struct {
		name     string
		input    []Employee
		expected []Employee
		sortBy   string
	}{
		{
			name: "Sort by LastName then FirstName then Salary descending",
			input: []Employee{
				{FirstName: "Alice", LastName: "Anderson", Age: 25, Salary: 50000},
				{FirstName: "Bob", LastName: "Brown", Age: 30, Salary: 60000},
				{FirstName: "Charlie", LastName: "Clark", Age: 22, Salary: 40000},
				{FirstName: "David", LastName: "Davis", Age: 27, Salary: 70000},
				{FirstName: "Emma", LastName: "Elliott", Age: 24, Salary: 55000},
			},
			expected: []Employee{
				{FirstName: "David", LastName: "Davis", Age: 27, Salary: 70000},
				{FirstName: "Bob", LastName: "Brown", Age: 30, Salary: 60000},
				{FirstName: "Emma", LastName: "Elliott", Age: 24, Salary: 55000},
				{FirstName: "Alice", LastName: "Anderson", Age: 25, Salary: 50000},
				{FirstName: "Charlie", LastName: "Clark", Age: 22, Salary: 40000},
			},
			sortBy: "LastNameFirstNameSalary",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := make([]Employee, len(tc.input))
			copy(got, tc.input)
			sortEmployees(got, tc.sortBy)

			for i, emp := range got {
				assert.Equal(t, tc.expected[i].FirstName, emp.FirstName, fmt.Sprintf("FirstName mismatch at index %d", i))
				assert.Equal(t, tc.expected[i].LastName, emp.LastName, fmt.Sprintf("LastName mismatch at index %d", i))
				assert.Equal(t, tc.expected[i].Age, emp.Age, fmt.Sprintf("Age mismatch at index %d", i))
				assert.Equal(t, tc.expected[i].Salary, emp.Salary, fmt.Sprintf("Salary mismatch at index %d", i))
			}
		})
	}
}

func sortEmployees(employees []Employee, sortBy string) {
	switch sortBy {
	case "LastNameFirstNameSalary":
		sort.Slice(employees, func(i, j int) bool {
			if employees[i].LastName != employees[j].LastName {
				return employees[i].LastName < employees[j].LastName
			}
			if employees[i].FirstName != employees[j].FirstName {
				return employees[i].FirstName < employees[j].FirstName
			}
			return employees[i].Salary > employees[j].Salary
		})
	}
}
