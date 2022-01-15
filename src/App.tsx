import React from 'react';
import './App.css';
import axios from "axios";
import { LocalizationProvider, DatePicker } from '@mui/lab';
import AdapterDateFns from '@mui/lab/AdapterDateFns';
import { Button, Card, Grid, List, ListItem, ListItemButton, TextField, Typography } from '@mui/material';
import { Box } from '@mui/system';
import DiaryAppBar from "./DiaryAppBar"
import Utils from "./Utils";

type lineType = { day: String; memo: String; };

/** App アプリケーションメインコンポーネント */
class App extends React.Component {

  /** カスタムstate */
  state: { date: Date | null, memo: string, items: lineType[] } = { date: null, memo: "", items: [] };

  /** コンストラクタ */
  constructor(props: any) {
    super(props);
    this.state = { date: new Date(), memo: "", items: [] };
  }

  /** マウントライフサイクルメソッド */
  componentDidMount() {
    this.showLines();
  }

  /** 画面レンダリング */
  render() {
    return (
      <div className="App">
        <DiaryAppBar />
        <Box sx={{ m: 2 }}>
          <Grid container direction="column" alignItems="center">
            <Grid item>
              <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker label="Date" value={this.state.date} onChange={this.changeDate}
                  inputFormat="yyyy/MM/dd"
                  mask="____/__/__"
                  renderInput={(params) => <TextField {...params} />} />
              </LocalizationProvider>
              <TextField label="Memo" value={this.state.memo} InputLabelProps={{ shrink: true }} onChange={this.changeMemo} />
              <Button variant="outlined" onClick={this.clickSend}>send</Button>
            </Grid>
          </Grid>
        </Box>
        <Card sx={{ m: 2 }}>
          <List>
            {this.state.items.map((v) => <ListItem><ListItemButton onClick={e => this.clickList(e, v)}>{v.day} {v.memo}</ListItemButton></ListItem>)}
          </List>
        </Card>
      </div>
    );
  }

  /** 日付変更イベント */
  changeDate = (newValue: Date | null) => { this.setState({ date: newValue }) }
  /** メモ変更イベント */
  changeMemo = (e: React.ChangeEvent<HTMLInputElement>) => this.setState({ memo: e.target.value })
  /** 送信 クリックイベント */
  clickSend = async () => {
    if (this.state.date != null) {
      let url = Utils.formatDate(this.state.date);
      url = `./api/${url}`;

      await axios.post(url, { Memo: this.state.memo });
      this.showLines();
    }
  }
  /** リストクリックイベント */
  clickList = (e: any, l: lineType) => {
    const y = Number(l.day.substring(0, 4))
    const m = Number(l.day.substring(5, 7)) - 1
    const d = Number(l.day.substring(8, 10))

    const day = new Date(y, m, d)
    this.setState({ date: day, memo: l.memo.trim() })
  }

  showLines = async () => {
    const url = `./api/${Utils.formatMonth(new Date())}`;
    const r = await axios.get(url);
    const s = this.state;
    s.items = r.data as lineType[];
    this.setState(s);
  }
}

export default App;
