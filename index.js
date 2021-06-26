const app = require(`express`)();
app.listen(3000,function(){
console.log("Siteye Bağlanıldı");
});
app.get("/", (req, res) => {
	res.sendFile(__dirname + "/index.html")
})