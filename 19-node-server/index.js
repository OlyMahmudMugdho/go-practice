const express = require('express');
const bodyParser = require('body-parser');
const cors = require("cors");


const app = express();

app.use(cors());
app.use(bodyParser.urlencoded({extended : false}));
app.use(bodyParser.json());


app.get("/get-end", (req,res) => {
    res.json({
        message : "hello world"
    })
})

app.post("/post-end", (req,res) => {
    const data = req.body;
    console.log(data)
    res.json(data)
} )


app.listen(3000, () => {
    console.info("server is running on port 3000")
})