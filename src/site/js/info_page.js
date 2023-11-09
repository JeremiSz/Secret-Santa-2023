{
    let wishlistItems = document.getElementsByClassName("wishlist-item");
    for (item of wishlistItems){
        let button = item.children[0]
        button.onclick = function(){
            button.parentElement.remove();
        };
    }
}
let wishlist = document.getElementById("wishlist");

{
    let addButton = document.getElementById("add-button");
    addButton.onclick = function(){
        let wishlistItem = createButton();
        wishlist.appendChild(wishlistItem);
    }
    
}

function createButton(name){
    let wishlistItemTemplate = document.createElement("div");
    wishlistItemTemplate.classList.add("singleInput","wishlist-item");
    let button = document.createElement("button");
    button.innerHTML = "X";
    button.onclick = function(){
        wishlistItemTemplate.remove();
    }
    let input = document.createElement("input");
    input.type = "text";
    input.name = "list"
    input.placeholder = "Enter item";
    wishlistItemTemplate.appendChild(button);
    wishlistItemTemplate.appendChild(document.createElement)
    wishlistItemTemplate.appendChild(input);
    return wishlistItemTemplate;
}