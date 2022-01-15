import { AppBar, Toolbar, Typography } from "@mui/material";
import React from "react";

export default class DiaryAppBar extends React.Component {
  render() {
    return (
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Diary
          </Typography>
        </Toolbar>
      </AppBar>
    )
  }
}