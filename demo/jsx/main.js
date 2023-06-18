let Button = ({ text }) =>
  <div className="button">{text}</div>

document.body.appendChild(Button({text: 'Click'}))
