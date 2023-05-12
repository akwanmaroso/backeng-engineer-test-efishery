import express, { Express, Request, Response } from "express";

const app: Express = express();
const port = 4004;

app.get("/", (req: Request, res: Response) => {
  res.send("oke");
});

app.listen(port, () => {
  console.log(`Server running on port ${port}`);
});
