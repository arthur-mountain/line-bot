import 'dotenv/config';
import express from 'express';
import notion from './modal/notion.js';
const app = express();

app.use(express.urlencoded({ extended: false }));
app.use(express.json());

app.get('/', async (_, res) => {
  const contents = await notion.getDataBaseContents();

  res.send({ statusCode: 200, data: contents });
});

app.post('/page', async (req, res) => {
  const blocks = await notion.getBlocksData({ ...req.body });

  res.send({ statusCode: 200, data: blocks });
});

app.listen(process.env.PORT)