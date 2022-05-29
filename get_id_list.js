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

function main(){
    const list = getIdList()
    console.log(`共 ${list.length} 条`)
    console.log(list.join())
}

main()
