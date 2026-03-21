CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255),
    descricao VARCHAR(255),
    preco DECIMAL(10, 2),
    quantidade INTEGER
);

INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES
('Teclado Mecânico', 'Switch Blue RGB', 250.00, 15),
('Mouse Gamer', '12000 DPI Sensor óptico', 120.50, 30),
('Monitor 24 Pol', 'Full HD 144Hz', 1100.00, 10),
('Headset 7.1', 'Cancelamento de ruído', 340.90, 20);

# docker build -t meu-postgres-loja .
# docker run -d --name pg-loja -p 5432:5432 meu-postgres-loja