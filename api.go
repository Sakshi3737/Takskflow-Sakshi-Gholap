func getProjects(c *gin.Context) {
 userID := c.GetString("user_id")

 rows, _ := db.Query(`
 SELECT DISTINCT p.id, p.name, p.description
 FROM projects p
 LEFT JOIN tasks t ON t.project_id = p.id
 WHERE p.owner_id=$1 OR t.assignee_id=$1`, userID)

 var projects []gin.H
 for rows.Next() {
  var id, name, desc string
  rows.Scan(&id, &name, &desc)

  projects = append(projects, gin.H{
   "id": id, "name": name, "description": desc,
  })
 }

 c.JSON(200, gin.H{"projects": projects})
}
