package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	guuid "github.com/google/uuid"
	"myfirstgosql/models"
)

type IProductRepository interface {
	Insert(product models.Product) (*models.Product, error)
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
	}
)

type ProductRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func NewProductRepository(db *sql.DB) IProductRepository {
	ps := make(map[string]*sql.Stmt, len(productQueries))
	for n, v := range productQueries {
		p, err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n] = p
	}
	return &ProductRepository{
		db, ps,
	}
}
func (r *ProductRepository) Insert(product models.Product) (*models.Product, error) {
	id := guuid.New()
	product.Id = id.String()
	res, err := r.ps["insertProduct"].Exec(product.Id, product.ProductCode, product.ProductName)
	if err != nil {
		return nil, err
	}

	affectedNo, err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &product, nil
}

func (r *ProductRepository) FindOneById(id string) (*models.Product, error) {
	rows, err := r.ps["productFindOneById"].Query(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	res := new(models.Product)
	err = rows.Scan(&res.Id, &res.ProductCode, &res.ProductName)
	if err != nil {
		panic(err)
	}
	return res, nil
}
func (r *ProductRepository) FindAllByNameLike(name string) ([]*models.Product, error) {
	rows, err := r.ps["productFindAllByNameLike"].Query(name)
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
	rows, err := r.ps["productFindAllProductPaging"].Query(pageNo, totalPerPage)
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
	row := r.db.QueryRow("select count(id) from m_product")
	res := new(models.TotalProduct)
	err := row.Scan(&res.Count)
	if err != nil {
		return -1, nil
	}
	return res.Count, nil
}
