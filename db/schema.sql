CREATE TABLE IF NOT EXISTS t_restaurant (
    id char(6) NOT NULL,
    _name text NOT NULL,
    _image text NOT NULL,
    _location point NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS t_product (
    id char(6) NOT NULL,
    restaurant_id char(6) NOT NULL,
    _category text NOT NULL,
    _name text NOT NULL,
    _description text NOT NULL,
    _image text,
    _price_value integer NOT NULL,
    _price_currency text NOT NULL,
    PRIMARY KEY(id)
);

INSERT INTO t_restaurant (id, _name, _image, _location)
VALUES (
    '111111',
    'La Bocatería',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    point(41.404564, 2.172751)
);

INSERT INTO t_restaurant (id, _name, _image, _location)
VALUES (
    '222222',
    'Las Hamburguesas',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    point(41.397740, 2.164425)
);

INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00001',
    '111111',
    'preferidos',
    'Bocata de atún',
    'blablablablablablalalbal vbakb kvbadk',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    300,
    'EUR'
);

INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00002',
    '111111',
    'bebidas',
    'Cocacola',
    'muy fría',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    150,
    'EUR'
);

INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00003',
    '111111',
    'veganos',
    'Bocadillo de aguacate',
    'muy vegano',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    500,
    'EUR'
);

INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00004',
    '222222',
    'clasicos',
    'Hamburguesa completa',
    'con todo',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    900,
    'EUR'
);

INSERT INTO t_product (
    id,
    restaurant_id,
    _category,
    _name,
    _description,
    _image,
    _price_value,
    _price_currency
) VALUES (
    'p00005',
    '222222',
    'clasicos',
    'Hamburguesa doble',
    'doble hamburguesa',
    'asesorexcelente.com/wp-content/uploads/2021/05/bocateria.png',
    1100,
    'EUR'
);