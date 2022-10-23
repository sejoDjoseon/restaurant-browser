CREATE TABLE IF NOT EXISTS restaurant (
    id char(6) NOT NULL,
    _name text NOT NULL,
    _image text NOT NULL,
    _location point NOT NULL,
    PRIMARY KEY(id)
);

INSERT INTO restaurant (id, _name, _image, _location)
VALUES (
    '111111',
    'La Bocatería',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    point(41.404564, 2.172751)
);

INSERT INTO restaurant (id, _name, _image, _location) 
VALUES (
    '222222',
    'Las Hamburguesas',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    point(41.397740, 2.164425)
)