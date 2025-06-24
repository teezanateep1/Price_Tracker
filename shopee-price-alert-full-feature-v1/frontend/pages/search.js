
import { useState } from 'react';

export default function Search() {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState([]);

  const search = async () => {
    const res = await fetch('/api/search?q=' + encodeURIComponent(query));
    const data = await res.json();
    setResults(data);
  };

  const follow = async (item) => {
    alert("ระบบติดตามสินค้านี้ยังเป็น mock ครับ: " + item.name);
    // TODO: เชื่อม API POST /api/subscribe พร้อมบันทึก itemId/shopId
  };

  return (
    <div style={{ padding: 20 }}>
      <h2>🔍 ค้นหาสินค้าบน Shopee</h2>
      <input value={query} onChange={(e) => setQuery(e.target.value)} placeholder="เช่น iPhone 15" />
      <button onClick={search}>ค้นหา</button>
      <div style={{ marginTop: 20 }}>
        {results.map((item, index) => (
          <div key={index} style={{ marginBottom: 20, border: '1px solid #ccc', padding: 10 }}>
            <img src={item.image} width="100" alt={item.name} />
            <h4>{item.name}</h4>
            <p>ราคา: {item.price} บาท</p>
            <button onClick={() => follow(item)}>ติดตาม</button>
            &nbsp;
            <a href={item.affiliate_url} target="_blank" rel="noreferrer">
              <button>ไปยัง Shopee</button>
            </a>
          </div>
        ))}
      </div>
    </div>
  );
}
