import 'bulma/css/bulma.css';
import { useEffect, useState } from 'react';
import axios from "axios";
import AppBar from "./MuAppBar"
import Utils from "./Utils";

type lineType = { Day: string; Memo: string; };
type listType = { WritedMonths: string[], Lines: lineType[] };

/** App アプリケーションメインコンポーネント */
export default function App() {
  const [date, setDate] = useState<Date | null>(new Date());
  const [memo, setMemo] = useState("");
  const [items, setItems] = useState<listType>({ WritedMonths: [], Lines: new Array<lineType>() });
  const [selectMonth, setSelectMonth] = useState("2022-02")

  /** 送信 クリックイベント */
  const sendButtonClick = async () => {
    if (date != null) {
      let url = Utils.formatDate(date);
      url = `./api/${url}`;

      await axios.post(url, { Memo: memo });
      showLines();
    }
  }

  /** リストクリックイベント */
  const listClick = (e: any, l: lineType) => {
    // const y = Number(l.Day.substring(0, 4))
    // const m = Number(l.Day.substring(5, 7)) - 1
    // const d = Number(l.Day.substring(8, 10))

    // const day = new Date(y, m, d)
    setDate2(l.Day)
    setMemo(l.Memo)
  }

  const showLines = async () => {
    const url = `./api/${Utils.formatMonth(new Date())}`;
    axios.get(url).then(r => {
      const i = r.data as listType;
      setItems(i);
    })
  }

  const setDate2 = (value: string) => {
    let e = value === "" ? null : new Date(value)

    setDate(e)
    if (e == null) {
      e = new Date()
    }

    var d = e.getFullYear() + "-" + (e.getMonth() + 1)
    var a = items?.WritedMonths?.filter((v, i) => v === d)

    if (a.length === 0 && items.WritedMonths.length > 0) {
      d = items.WritedMonths[0]
      setSelectMonth(d)
    } else {
      setSelectMonth(d)
    }
  }

  useEffect(() => { showLines() }, [])

  return (
    <div>
      <AppBar />

      <div className="card px-10">
        <div className="card-header">
          <input type="date" onChange={e => setDate2(e.target.value)} />
          <input type="text" value={memo} onChange={e => setMemo(e.target.value)} />
          <button className='button' onClick={sendButtonClick}>send</button>
        </div>
        <div className="card-content">
          <div className="select">
            <select value={selectMonth} onChange={e => setSelectMonth(e.target.value)}>
              {items.WritedMonths.map((v) => <option value={v}>{v}</option>)}
            </select>
            aj
          </div>
          <table>
            <tbody>
              {items.Lines.map((v) => <tr onClick={e => listClick(e, v)}><td>{v.Day} {v.Memo}</td></tr>)}
            </tbody>
          </table>
        </div>
      </div>
    </div>

  )
}