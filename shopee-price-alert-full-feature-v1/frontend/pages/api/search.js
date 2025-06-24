export default async function handler(req, res) {
  const { q } = req.query;
  const response = await fetch(`http://localhost:8080/api/search?q=${q}`);
  const data = await response.json();
  res.status(200).json(data);
}
