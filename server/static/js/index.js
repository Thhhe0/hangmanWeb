const btnSb = document.querySelector('.showSb')
const bg = document.querySelector('.bg')
const sbCard = document.querySelector('.scoreboard')
const sbList = document.querySelectorAll('.scoreboard-sorted > tbody > *')
const sbBtnE = document.querySelector('.scoreboard > nav > .btn-easy')
const sbBtnN = document.querySelector('.scoreboard > nav > .btn-normal')
const sbBtnH = document.querySelector('.scoreboard > nav > .btn-hard')
const utils = {
  difficulty: [
    { key: 1, value: sbBtnE },
    { key: 2, value: sbBtnN },
    { key: 3, value: sbBtnH },
  ]
}

btnSb.addEventListener('click', (e) => {
  bg.classList.remove('none')
  sbCard.classList.remove('none')
})

bg.addEventListener('click', (e) => {
  bg.classList.add('none')
  sbCard.classList.add('none')
})

sbBtnE.addEventListener('click', (e) => {
  displaySb(1)
  setActiveBtn(1)
})

sbBtnN.addEventListener('click', (e) => {
  displaySb(2)
  setActiveBtn(2)
})

sbBtnH.addEventListener('click', (e) => {
  displaySb(3)
  setActiveBtn(3)
})

const displaySb = (diff) => {
  sbList.forEach((v) => {
    if (Number(v.getAttribute('data-difficulty')) != diff) {
      v.classList.add('none')
    } else {
      v.classList.remove('none')
    }
  })
}

const setActiveBtn = (diff) => {
  utils.difficulty.forEach((v) => {
    if (diff != v.key) {
      v.value.classList.remove('btn-sb-active')
    } else {
      v.value.classList.add('btn-sb-active')
    }
  })
}

(function () {
  displaySb(3)
  setActiveBtn(3)
})()
