import { useEffect, useState } from 'react';
import './App.css';
import axios from "axios";
import { LocalizationProvider, DatePicker } from '@mui/lab';
import AdapterDateFns from '@mui/lab/AdapterDateFns';
import { Alert, Button, Card, Grid, List, ListItem, ListItemButton, Snackbar, TextField, Typography } from '@mui/material';
import { Box } from '@mui/system';
import DiaryAppBar from "./DiaryAppBar"
import Utils from "./Utils";

type lineType = { day: string; memo: string; };

/** App アプリケーションメインコンポーネント */
function App() {

  const [date, setDate] = useState<Date | null>(new Date());
  const [memo, setMemo] = useState("");
  const [items, setItems] = useState(new Array<lineType>());

  /** 送信 クリックイベント */
  const clickSend = async () => {
    if (date != null) {
      let url = Utils.formatDate(date);
      url = `./api/${url}`;

      await axios.post(url, { Memo: memo });
      showLines();
    }
  }

  /** リストクリックイベント */
  const clickList = (e: any, l: lineType) => {
    const y = Number(l.day.substring(0, 4))
    const m = Number(l.day.substring(5, 7)) - 1
    const d = Number(l.day.substring(8, 10))

    const day = new Date(y, m, d)
    setDate(day)
    const memo = l.memo
    setMemo(memo)
  }

  const showLines = async () => {
    const url = `./api/${Utils.formatMonth(new Date())}`;
    axios.get(url).then(r => {
      const i = r.data as [];
      setItems(i);
    })
  }

  useEffect(() => { showLines() }, [])

  return (
    <div className="App">
      <DiaryAppBar />
      <Box sx={{ m: 2 }}>
        <Grid container direction="column" alignItems="center">
          <Grid item>
            <LocalizationProvider dateAdapter={AdapterDateFns}>
              <DatePicker label="Date" value={date} onChange={e => setDate(e)}
                inputFormat="yyyy/MM/dd"
                mask="____/__/__"
                renderInput={(params) => <TextField {...params} />} />
            </LocalizationProvider>
            <TextField label="Memo" value={memo} InputLabelProps={{ shrink: true }} onChange={e => setMemo(e.target.value)} />
            <Button variant="outlined" onClick={clickSend}>send</Button>
          </Grid>
        </Grid>
      </Box>
      <Card sx={{ m: 2 }}>
        <List>
          {items.map((v) => <ListItem><ListItemButton onClick={e => clickList(e, v)}>{v.day} {v.memo}</ListItemButton></ListItem>)}
        </List>
      </Card>
    </div>)
}

export default App;
