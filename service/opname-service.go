package service

import(
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"serviceOpname-v2/controller/dto"
	"serviceOpname-v2/config/entity"
	"serviceOpname-v2/config/entity/helper"
	"serviceOpname-v2/repository"
)

type OpnameService interface {
	Update(b dto.OpnameUpdDTO) entity.Opname
	All() []entity.Opname
	FindById(opnameID uint64) entity.Opname
	IsAllowedToEdit(userID string, opnameID uint64) bool
	GetPaginate(param helper.Pagination) (repository.RepositoryResult, int)
}

type opnameService struct {
	opnameRepository repository.OpnameRepository
}

func NewOpnameService(opnameRepo repository.OpnameRepository) OpnameService {
	return &opnameService{
		opnameRepository: opnameRepo,
	}
}

func (service *opnameService) Update(b dto.OpnameUpdDTO) entity.Opname {
	opname := entity.Opname{}
	err := smapping.FillStruct(&opname, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v :", err)
	}
	res := service.opnameRepository.UpdateOpname(opname)
	return res
}

func (service *opnameService) All() []entity.Opname {
	return service.opnameRepository.AllOpname()
}

func (service *opnameService) FindById(opnameID uint64) entity.Opname {
	return service.opnameRepository.FindOpnameByID(opnameID)
}

func (service *opnameService) IsAllowedToEdit(userID string, opnameID uint64) bool {
	b := service.opnameRepository.FindOpnameByID(opnameID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}

func (service *opnameService) GetPaginate(pagination helper.Pagination) (repository.RepositoryResult, int){
	return service.opnameRepository.Pagination(pagination)
}