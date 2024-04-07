package practices

import "strings"

/*
Write a function that receives an array of strings that represent friend connections along with the names of two people and returns a number representing the degrees of separation between the two people.
The connections will be represented by an array of strings with each string taking the format name1:name2 for examples alice:bob.
You can assume that the strings representing the connections will always be lower case a-z only.
The names of the people to find the degrees of separation between will always be non- empty strings e.g. "alice" or "bob".
Your function will return the number of degrees of separation between the two people. If no connection can be made through friends or friends of friends etc then return -1. Example 10
connections = ["fred:joe", "joe:mary", "mary:fred","mary:bill"]
name1 = "fred"
name2 = "bill"
Result = 2
The expected result is 2 because fred is friends with mary, and mary is friends with bill. That is, bill is of distance 2 from fred. O
fred > mary > bill

Example 2
connections = ["fred:joe", "joe:mary", "kate:sean", "sean:sally"]
name1 = kate
name2 = sally
=> 1
*/

func friendCircle(conn []string, name1, name2 string) int {
	connMap := make(map[string][]string)
	for _, v := range conn {
		names := strings.Split(v, ":")
		n1, n2 := names[0], names[1]
		connMap[n1] = append(connMap[n1], n2)
		connMap[n2] = append(connMap[n2], n1)
	}

	// BFS
	queue := []string{name1}
	visited := make(map[string]bool)
	visited[name1] = true

	distance := 0
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		for _, v := range connMap[item] {
			if !visited[v] {
				if v == name2 {
					return distance
				} else {
					queue = append(queue, v)
					visited[v] = true
				}
			}
		}
		distance++
	}

	return -1
}
