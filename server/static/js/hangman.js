const input = document.querySelector('.letter')
const inputNick = document.querySelector('.nick')
const form = document.querySelector('.letter-form')

if (!inGame) {
  inputNick.focus()
  form.addEventListener('submit', (e) => e.preventDefault())
} else {
  input.focus()
}
