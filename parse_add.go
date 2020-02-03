package jint

import "strconv"

func (p *parse) AddKeyValue(key string, newVal []byte, path ...string) error {
	lenp := len(path)
	lenv := len(key)
	var curr *node
	var err error
	if lenv == 0 {
		return NULL_KEY_ERROR()
	}
	json, err := p.Get(path...)
	if err != nil {
		return err
	}
	curr = p.core
	if lenp == 0 {
		for _, d := range curr.down {
			if d.label == key {
				return KEY_ALREADY_EXISTS_ERROR()
			}
		}
		if len(json) >= 2 {
			if json[0] == 123 && json[len(p.json)-1] == 125 {
				newKV := []byte(`,"` + key + `":` + string(newVal))
				if lenv >= 2 {
					if newVal[0] == 91 || newVal[0] == 123 {
						newNode := CreateNode(nil)
						pCore(newVal, newNode)
						newNode.label = key
						newNode.value = newVal
						if len(p.json) == 2 {
							p.json = replace(p.json, newKV[1:], len(p.json)-1, len(p.json)-1)
						} else {
							p.json = replace(p.json, newKV, len(p.json)-1, len(p.json)-1)
						}
						curr.down = append(curr.down, newNode)
						newNode.up = curr
						return nil
					}
				}
				p.json = replace(p.json, newKV, len(p.json)-1, len(p.json)-1)
				newNode := CreateNode(nil)
				newNode.label = key
				newNode.value = newVal
				curr.down = append(curr.down, newNode)
				newNode.up = curr
				return nil
			}
			return OBJECT_EXPECTED_ERROR()
		}
		return BAD_JSON_ERROR(0)
	}
	for _, d := range curr.up.down {
		if d.label == key {
			return KEY_ALREADY_EXISTS_ERROR()
		}
	}
	if len(json) >= 2 {
		if json[0] == 123 && json[len(p.json)-1] == 125 {
			if lenv >= 2 {
				if newVal[0] == 91 || newVal[0] == 123 {
					newNode := CreateNode(nil)
					pCore(newVal, newNode)
					newNode.label = key
					newNode.value = newVal
					curr.down = append(curr.down, newNode)
					newNode.up = curr
					p.json, err = AddKeyValue(p.json, key, newVal, path...)
					if err != nil {
						return err
					}
					for i := 0; i < lenp; i++ {
						newNode = newNode.up
						newNode.value, err = Get(p.json, path[:lenp-i]...)
						if err != nil {
							return err
						}
					}
					return nil
				}
			}
			newNode := CreateNode(nil)
			newNode.label = key
			newNode.value = newVal
			curr.down = append(curr.down, newNode)
			newNode.up = curr
			p.json, err = AddKeyValue(p.json, key, newVal, path...)
			if err != nil {
				return err
			}
			for i := 0; i < lenp; i++ {
				newNode = newNode.up
				newNode.value, err = Get(p.json, path[:lenp-i]...)
				if err != nil {
					return err
				}
			}
			return nil
		}
		return OBJECT_EXPECTED_ERROR()
	}
	return BAD_JSON_ERROR(0)
}

func (p *parse) Add(newVal []byte, path ...string) error {
	lenp := len(path)
	lenv := len(newVal)
	var curr *node
	var err error
	if lenv == 0 {
		return NULL_KEY_ERROR()
	}
	json, err := p.Get(path...)
	if err != nil {
		return err
	}
	curr = p.core
	if lenp == 0 {
		if len(json) >= 2 {
			if json[0] == 91 && json[len(p.json)-1] == 93 {
				newValue := []byte(`,` + string(newVal))
				if lenv >= 2 {
					if newVal[0] == 91 || newVal[0] == 123 {
						newNode := CreateNode(nil)
						pCore(newVal, newNode)
						index := len(curr.down)
						newNode.label = strconv.Itoa(index)
						newNode.value = newVal
						if len(p.json) == 2 {
							p.json = replace(p.json, newValue[1:], len(p.json)-1, len(p.json)-1)
						} else {
							p.json = replace(p.json, newValue, len(p.json)-1, len(p.json)-1)
						}
						curr.down = append(curr.down, newNode)
						newNode.up = curr
						return nil
					}
				}
				p.json = replace(p.json, newValue, len(p.json)-1, len(p.json)-1)
				newNode := CreateNode(nil)
				index := len(curr.down)
				newNode.label = strconv.Itoa(index)
				newNode.value = newVal
				curr.down = append(curr.down, newNode)
				newNode.up = curr
				return nil
			}
			return ARRAY_EXPECTED_ERROR()
		}
		return BAD_JSON_ERROR(0)
	}
	if len(json) >= 2 {
		if json[0] == 91 && json[len(json)-1] == 93 {
			if lenv >= 2 {
				if newVal[0] == 91 || newVal[0] == 123 {
					newNode := CreateNode(nil)
					pCore(newVal, newNode)
					index := len(curr.down)
					newNode.label = strconv.Itoa(index)
					newNode.value = newVal
					curr.down = append(curr.down, newNode)
					newNode.up = curr
					p.json, err = Add(p.json, newVal, path...)
					if err != nil {
						return err
					}
					for i := 0; i < lenp; i++ {
						newNode = newNode.up
						newNode.value, err = Get(p.json, path[:lenp-i]...)
						if err != nil {
							return err
						}
					}
					return nil
				}
			}
			newNode := CreateNode(nil)
			index := len(curr.down)
			newNode.label = strconv.Itoa(index)
			newNode.value = newVal
			curr.down = append(curr.down, newNode)
			newNode.up = curr
			p.json, err = Add(p.json, newVal, path...)
			if err != nil {
				return err
			}
			for i := 0; i < lenp; i++ {
				newNode = newNode.up
				newNode.value, err = Get(p.json, path[:lenp-i]...)
				if err != nil {
					return err
				}
			}
			return nil
		}
		return ARRAY_EXPECTED_ERROR()
	}
	return BAD_JSON_ERROR(0)
}

func (p *parse) Insert(newVal []byte, newIndex int, path ...string) error {
	lenp := len(path)
	lenv := len(newVal)
	var curr *node
	var err error
	if lenv == 0 {
		return NULL_KEY_ERROR()
	}
	json, err := p.Get(path...)
	if err != nil {
		return err
	}
	curr = p.core
	if lenp == 0 {
		if len(json) >= 2 {
			if json[0] == 91 && json[len(p.json)-1] == 93 {
				if lenv >= 2 {
					if newVal[0] == 91 || newVal[0] == 123 {
						newNode := CreateNode(nil)
						pCore(newVal, newNode)
						err = newNode.insert(curr, newIndex)
						if err != nil {
							return err
						}
						newNode.value = newVal
						p.json, err = Insert(p.json, newIndex, newVal, path...)
						curr.down = append(curr.down, newNode)
						newNode.up = curr
						return nil
					}
				}
				p.json, err = Insert(p.json, newIndex, newVal, path...)
				newNode := CreateNode(nil)
				err = newNode.insert(curr, newIndex)
				if err != nil {
					return err
				}
				newNode.value = newVal
				curr.down = append(curr.down, newNode)
				newNode.up = curr
				return nil
			}
			return ARRAY_EXPECTED_ERROR()
		}
		return BAD_JSON_ERROR(0)
	}
	if len(json) >= 2 {
		if json[0] == 91 && json[len(json)-1] == 93 {
			if lenv >= 2 {
				if newVal[0] == 91 || newVal[0] == 123 {
					newNode := CreateNode(nil)
					pCore(newVal, newNode)
					err = newNode.insert(curr, newIndex)
					if err != nil {
						return err
					}
					newNode.value = newVal
					curr.down = append(curr.down, newNode)
					newNode.up = curr
					p.json, err = Insert(p.json, newIndex, newVal, path...)
					if err != nil {
						return err
					}
					for i := 0; i < lenp; i++ {
						newNode = newNode.up
						newNode.value, err = Get(p.json, path[:lenp-i]...)
						if err != nil {
							return err
						}
					}
					return nil
				}
			}
			newNode := CreateNode(nil)
			err = newNode.insert(curr, newIndex)
			if err != nil {
				return err
			}
			newNode.value = newVal
			curr.down = append(curr.down, newNode)
			newNode.up = curr
			p.json, err = Insert(p.json, newIndex, newVal, path...)
			if err != nil {
				return err
			}
			for i := 0; i < lenp; i++ {
				newNode = newNode.up
				newNode.value, err = Get(p.json, path[:lenp-i]...)
				if err != nil {
					return err
				}
			}
			return nil
		}
		return ARRAY_EXPECTED_ERROR()
	}
	return BAD_JSON_ERROR(0)
}