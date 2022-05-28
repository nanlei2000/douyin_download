// @ts-check

function getIdList(){
    const eles  = document.getElementsByClassName("ECMy_Zdt")
    const idList = []
    for (const ele of eles) {
        /** @type {HTMLAnchorElement} */
        // @ts-expect-error
        const a = ele.firstChild
        const part =  a.href.split('/')
        idList.push(part[part.length-1])
    }

    return idList
}

console.log(getIdList().join())