package main

import (
	"reflect"
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

type employeeSorter struct {
	employees []Employee
}

func (s employeeSorter) Len() int {
	return len(s.employees)
}

func (s employeeSorter) Swap(i, j int) {
	s.employees[i], s.employees[j] = s.employees[j], s.employees[i]
}

func (s employeeSorter) Less(i, j int) bool {
	if s.employees[i].LastName != s.employees[j].LastName {
		return s.employees[i].LastName < s.employees[j].LastName
	}
	if s.employees[i].FirstName != s.employees[j].FirstName {
		return s.employees[i].FirstName < s.employees[j].FirstName
	}
	return s.employees[i].Salary > s.employees[j].Salary
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
			var sortedEmployees employeeSorter = employeeSorter{employees: tc.employees}
			sort.Sort(sortedEmployees)

			for i, employee := range sortedEmployees.employees {
				assert.Equal(t, tc.expected[i].FirstName, employee.FirstName, "FirstName is incorrect")
				assert.Equal(t, tc.expected[i].LastName, employee.LastName, "LastName is incorrect")
				assert.Equal(t, tc.expected[i].Age, employee.Age, "Age is incorrect")
				assert.Equal(t, tc.expected[i].Salary, employee.Salary, "Salary is incorrect")
			}
		})
	}
}
