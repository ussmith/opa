package pigs

huff_puffable[msg] {
    some c
    buildings[c]

    material := c.material

    not startswith(material, "bricks")
    msg := sprintf("Material '%v' is a huff_puffable", [material])
}

buildings[c] {
   some i
   c := input.neighborhood.houses[i]
}


buildings[c] {
   some i
   c := input.neighborhood.stores[i]
}
