
import Link from 'next/link';
import { useState } from 'react';

export default function Home() {
  const [query, setQuery] = useState('');

  const handleSearch = () => {
    if (query.trim()) {
      window.location.href = '/search?q=' + encodeURIComponent(query);
    }
  };

  return (
    <div style={{ padding: 30 }}>
      <h1>🔔 Shopee Price Tracker</h1>
      <p>ระบบติดตามราคาสินค้าที่คุณสนใจ พร้อมแจ้งเตือนเมื่อราคาลด ผ่าน LINE Notify หรือในเว็บของเรา</p>

      <div style={{ marginTop: 20 }}>
        <input
          placeholder="ค้นหาสินค้า เช่น iPhone 15"
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          style={{ padding: '8px', width: '300px' }}
        />
        <button onClick={handleSearch} style={{ marginLeft: 10 }}>ค้นหา</button>
      </div>

      <div style={{ marginTop: 40 }}>
        <h3>🔗 เมนูด่วน</h3>
        <ul>
          <li><Link href="/search">ค้นหาสินค้า</Link></li>
          <li><Link href="/subscriptions">สินค้าที่ติดตาม</Link></li>
          <li><Link href="/notifications">การแจ้งเตือนของฉัน</Link></li>
          <li><a href="https://notify-bot.line.me/oauth/authorize?response_type=code&client_id=YOUR_CLIENT_ID&redirect_uri=YOUR_REDIRECT_URI&scope=notify&state=xyz" target="_blank" rel="noreferrer">เชื่อมต่อ LINE Notify</a></li>
        </ul>
      </div>

      <div style={{ marginTop: 40 }}>
        <p><i>เข้าสู่ระบบด้วย Google เพื่อเริ่มต้นใช้งาน</i></p>
      </div>
    </div>
  );
}
