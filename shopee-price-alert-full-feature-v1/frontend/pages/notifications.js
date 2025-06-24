
import { useEffect, useState } from 'react';

export default function Notifications() {
  const [notifs, setNotifs] = useState([]);

  useEffect(() => {
    fetch('/api/notifications')
      .then(res => res.json())
      .then(setNotifs);
  }, []);

  const markRead = async (id) => {
    await fetch('/api/notifications/' + id + '/read', { method: 'POST' });
    setNotifs(notifs.map(n => n.id === id ? { ...n, is_read: true } : n));
  };

  return (
    <div style={{ padding: 20 }}>
      <h2>🔔 การแจ้งเตือนของคุณ</h2>
      {notifs.length === 0 && <p>ยังไม่มีการแจ้งเตือน</p>}
      {notifs.map(n => (
        <div key={n.id} style={{
          padding: 10,
          margin: '10px 0',
          border: '1px solid #ccc',
          background: n.is_read ? '#f0f0f0' : '#fff'
        }}>
          <b>{n.product_name}</b><br />
          {n.message}<br />
          <small>{new Date(n.created_at).toLocaleString()}</small><br />
          {!n.is_read && <button onClick={() => markRead(n.id)}>มาร์คว่าอ่านแล้ว</button>}
        </div>
      ))}
    </div>
  );
}
