import 'bulma/css/bulma.css';

const MuAppBar = () => {

  const burgerClick = () => {
    document.querySelector("#menu")?.classList.toggle('is-active');
    document.querySelector('.navbar-burger')?.classList.toggle('is-active');
  }

  return (
    <nav className="navbar is-dark" role="navigation" aria-label="main navigation">
      <div className="navbar-brand">
        <a className="navbar-item" href="./">la-diary</a>

        <button className="navbar-burger" onClick={burgerClick} aria-label="menu" aria-expanded="false">
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </button>
      </div>

      <div id="menu" className="navbar-menu">
        <div className="navbar-start">
        </div>

        <div className="navbar-end">
        </div>
      </div>

    </nav>
  )
}
export default MuAppBar
