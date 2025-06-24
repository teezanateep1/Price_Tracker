
import { useEffect, useState } from 'react';

export default function Subscriptions() {
  const [subs, setSubs] = useState([]);

  useEffect(() => {
    fetch('/api/subscriptions')
      .then(res => res.json())
      .then(setSubs);
  }, []);

  const updateAlert = async (id, type, threshold) => {
    await fetch('/api/subscriptions/' + id, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ alert_type: type, alert_threshold: threshold })
    });
    setSubs(subs.map(s => s.id === id ? { ...s, alert_type: type, alert_threshold: threshold } : s));
  };

  return (
    <div style={{ padding: 20 }}>
      <h2>📦 สินค้าที่คุณติดตาม</h2>
      {subs.length === 0 && <p>คุณยังไม่ได้ติดตามสินค้าใด ๆ</p>}
      {subs.map(s => (
        <div key={s.id} style={{
          padding: 10,
          margin: '10px 0',
          border: '1px solid #ccc'
        }}>
          <b>{s.product_name}</b><br />
          ราคาปัจจุบัน: {s.current_price} บาท<br />
          แจ้งเตือนเมื่อ:
          <select value={s.alert_type} onChange={(e) => updateAlert(s.id, e.target.value, s.alert_threshold)}>
            <option value="any_change">มีการเปลี่ยนแปลงราคา</option>
            <option value="percent">ลดเกินเปอร์เซ็นต์ที่กำหนด</option>
          </select>
          {s.alert_type === 'percent' && (
            <>
              &nbsp;
              <input type="number" value={s.alert_threshold}
                     onChange={e => updateAlert(s.id, 'percent', parseInt(e.target.value) || 10)}
                     style={{ width: 50 }} />%
            </>
          )}
        </div>
      ))}
    </div>
  );
}
