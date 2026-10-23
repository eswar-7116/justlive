package drop

func (d *Drop) Update(dt float32) {
	d.Life -= dt
	if d.Life <= 0 {
		d.Active = false
	}
}
