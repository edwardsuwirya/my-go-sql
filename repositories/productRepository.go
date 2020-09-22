package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	guuid "github.com/google/uuid"
	"myfirstgosql/config"
	"myfirstgosql/models"
)

type IProductRepository interface {
	Insert(product models.Product) (*models.Product, error)
	InsertPrice(productPrice models.ProductPrice) (*models.ProductPrice, error)
	InsertProductWithPrice(productWithPrice models.ProductWithPrice) (*models.ProductWithPrice, error)
	FindOneById(id string) (*models.Product, error)
	FindAllByNameLike(name string) ([]*models.Product, error)
	FindAllProductPaging(pageNo, totalPerPage int) ([]*models.Product, error)
	Count() (int64, error)
}

var (
	productQueries = map[string]string{
		"productFindOneById":          "select id,product_code,product_name from m_product where id=?",
		"productFindAllByNameLike":    "select id,product_code,product_name from m_product where product_name like ?",
		"productFindAllProductPaging": "select id,product_code,product_name from m_product order by id limit ?,?",
		"insertProduct":               "insert into m_product(id,product_code,product_name) values(?,?,?)",
		"insertProductPrice":          "insert into m_product_price(product_price_id,product_id,product_price,is_active) values(?,?,?,?)",
	}
)

type ProductRepository struct {
	session *config.Session
	ps      map[string]*sql.Stmt
}

func NewProductRepository(sf *config.SessionFactory) IProductRepository {
	ps := make(map[string]*sql.Stmt, len(productQueries))
	prodRepo := &ProductRepository{
		sf.GetSession(), ps,
	}
	for n, v := range productQueries {
		p, err := prodRepo.session.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n] = p
	}
	return prodRepo
}

func (r *ProductRepository) InsertProductWithPrice(productWithPrice models.ProductWithPrice) (*models.ProductWithPrice, error) {
	err := r.session.Begin()
	if err != nil {
		panic(err)
	}
	prodId := guuid.New().String()
	priceId := guuid.New().String()
	_, err = r.session.ExecStatement(r.ps["insertProduct"], prodId, productWithPrice.ProductCode, productWithPrice.ProductName)
	if err != nil {
		r.session.Rollback()
		return nil, err
	}
	_, err = r.session.ExecStatement(r.ps["insertProductPrice"], priceId, prodId, productWithPrice.Price, "0")
	if err != nil {
		r.session.Rollback()
		return nil, err
	}

	err = r.session.Commit()
	if err != nil {
		panic(err)
	}
	productWithPrice.Id = prodId
	productWithPrice.PriceId = prodId
	return &productWithPrice, nil

}
func (r *ProductRepository) Insert(product models.Product) (*models.Product, error) {
	id := guuid.New()
	product.Id = id.String()
	res, err := r.session.ExecStatement(r.ps["insertProduct"], product.Id, product.ProductCode, product.ProductName)
	if err != nil {
		return nil, err
	}

	affectedNo, err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &product, nil
}

func (r *ProductRepository) InsertPrice(productPrice models.ProductPrice) (*models.ProductPrice, error) {
	id := guuid.New()
	productPrice.PriceId = id.String()
	res, err := r.session.ExecStatement(r.ps["insertProductPrice"], productPrice.PriceId, productPrice.ProductId, productPrice.Price, "0")
	if err != nil {
		return nil, err
	}
	affectedNo, err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &productPrice, nil
}

func (r *ProductRepository) FindOneById(id string) (*models.Product, error) {
	row := r.session.QueryRowStatement(r.ps["productFindOneById"], id)
	res := new(models.Product)
	err := row.Scan(&res.Id, &res.ProductCode, &res.ProductName)
	if err != nil {
		panic(err)
	}
	return res, nil
}
func (r *ProductRepository) FindAllByNameLike(name string) ([]*models.Product, error) {
	rows, err := r.session.QueryStatement(r.ps["productFindAllByNameLike"], name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.Product, 0)
	for rows.Next() {
		res := new(models.Product)
		err = rows.Scan(&res.Id, &res.ProductCode, &res.ProductName)
		if err != nil {
			panic(err)
		}
		result = append(result, res)
	}
	return result, nil
}
func (r *ProductRepository) FindAllProductPaging(pageNo, totalPerPage int) ([]*models.Product, error) {
	rows, err := r.session.QueryStatement(r.ps["productFindAllProductPaging"], pageNo, totalPerPage)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]*models.Product, 0)
	for rows.Next() {
		res := new(models.Product)
		err = rows.Scan(&res.Id, &res.ProductCode, &res.ProductName)
		if err != nil {
			panic(err)
		}
		result = append(result, res)
	}
	return result, nil
}
func (r *ProductRepository) Count() (int64, error) {
	row := r.session.QueryRow("select count(id) from m_product")
	res := new(models.TotalProduct)
	err := row.Scan(&res.Count)
	if err != nil {
		return -1, nil
	}
	return res.Count, nil
}
