func getTasks(c *gin.Context) {
 projectID := c.Param("id")
 status := c.Query("status")
 assignee := c.Query("assignee")

 query := "SELECT id,title,status,priority FROM tasks WHERE project_id=$1"
 args := []interface{}{projectID}

 if status != "" {
  query += " AND status=$2"
  args = append(args, status)
 }

 rows, _ := db.Query(query, args...)

 var tasks []gin.H
 for rows.Next() {
  var id, title, status, priority string
  rows.Scan(&id, &title, &status, &priority)

  tasks = append(tasks, gin.H{
   "id": id,
   "title": title,
   "status": status,
   "priority": priority,
  })
 }

 c.JSON(200, gin.H{"tasks": tasks})
}
