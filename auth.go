func register(c *gin.Context) {
 var u struct {
  Name, Email, Password string
 }

 if err := c.BindJSON(&u); err != nil {
  c.JSON(400, gin.H{"error": "validation failed"})
  return
 }

 hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 12)

 _, err := db.Exec(
  "INSERT INTO users(name,email,password) VALUES($1,$2,$3)",
  u.Name, u.Email, string(hash),
 )

 if err != nil {
  c.JSON(400, gin.H{"error": "email already exists"})
  return
 }

 c.JSON(201, gin.H{"message": "registered"})
}

func login(c *gin.Context) {
 var u struct {
  Email, Password string
 }

 c.BindJSON(&u)

 var id, hash string
 err := db.QueryRow(
  "SELECT id,password FROM users WHERE email=$1",
  u.Email,
 ).Scan(&id, &hash)

 if err != nil || bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.Password)) != nil {
  c.JSON(401, gin.H{"error": "unauthorized"})
  return
 }

 exp := time.Now().Add(24 * time.Hour)

 token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
  UserID: id,
  Email: u.Email,
  RegisteredClaims: jwt.RegisteredClaims{
   ExpiresAt: jwt.NewNumericDate(exp),
  },
 })

 t, _ := token.SignedString(jwtKey)

 c.JSON(200, gin.H{"token": t})
}
