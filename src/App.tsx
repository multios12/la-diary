import { useEffect, useState } from 'react';
import './App.css';
import axios from "axios";
import { LocalizationProvider, DatePicker } from '@mui/lab';
import AdapterDateFns from '@mui/lab/AdapterDateFns';
import { Button, Card, createTheme, Grid, List, ListItem, ListItemButton, MenuItem, Select, TextField, ThemeProvider } from '@mui/material';
import AppBar from "./MuAppBar"
import Utils from "./Utils";

type lineType = { Day: string; Memo: string; };
type listType = { WritedMonths: [], Lines: lineType[] };

/** App アプリケーションメインコンポーネント */
export default function App() {

  const theme = createTheme({ palette: { mode: 'dark' } })
  const [date, setDate] = useState<Date | null>(new Date());
  const [memo, setMemo] = useState("");
  const [items, setItems] = useState({ WritedMonths: [], Lines: new Array<lineType>() });
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
    const y = Number(l.Day.substring(0, 4))
    const m = Number(l.Day.substring(5, 7)) - 1
    const d = Number(l.Day.substring(8, 10))

    const day = new Date(y, m, d)
    setDate2(day)
    setMemo(l.Memo)
  }

  const showLines = async () => {
    const url = `./api/${Utils.formatMonth(new Date())}`;
    axios.get(url).then(r => {
      const i = r.data as listType;
      setItems(i);
    })
  }

  const setDate2 = (e: Date | null) => {
    setDate(e)
    if (e == null) {
      e = new Date()
    }

    var d = e.getFullYear() + "-" + (e.getMonth() + 1)
    var a = items?.WritedMonths?.filter((v,i) => v===d)

    if(a.length == 0 && items.WritedMonths.length > 0) {
      setSelectMonth(items.WritedMonths[0])
    } else {
      setSelectMonth(d)
    }
  }

  useEffect(() => { showLines() }, [])

  return (
    <ThemeProvider theme={theme}>
      <div className="App">
        <AppBar />
        <Card sx={{ m: 1 }}>
          <Grid sx={{ m: 1 }} container direction="column" alignItems="center">
            <Grid item>
              <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker label="Date" value={date} onChange={e => setDate2(e)}
                  inputFormat="yyyy/MM/dd"
                  mask="____/__/__"
                  renderInput={(params) => <TextField {...params} />} />
              </LocalizationProvider>
              <TextField label="Memo" value={memo} InputLabelProps={{ shrink: true }} onChange={e => setMemo(e.target.value)} />
              <Button variant="outlined" onClick={sendButtonClick}>send</Button>
            </Grid>
          </Grid>
        </Card>
        <Card sx={{ m: 1 }}>
          <Select label="month" value={selectMonth} onChange={e => setSelectMonth(e.target.value)}>
            {items.WritedMonths.map((v) => <MenuItem value={v}>{v}</MenuItem>)}
          </Select>
          <List>
            {items.Lines.map((v) => <ListItem><ListItemButton onClick={e => listClick(e, v)}>{v.Day} {v.Memo}</ListItemButton></ListItem>)}
          </List>
        </Card>
      </div>
    </ThemeProvider>)
}