package handlers

import (
	"grpc-microservice-example/models"
	pb "grpc-microservice-example/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinHanlder struct {
	client pb.MovieServiceClient
}

func New(client pb.MovieServiceClient) *GinHanlder {
	return &GinHanlder{client: client}
}

func (g *GinHanlder) POST(ctx *gin.Context) {
	var movie models.Movies

	err := ctx.ShouldBind(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	data := &pb.Movie{
		Title: movie.Title,
		Genre: movie.Genre,
	}
	res, err := g.client.CreateMovie(ctx, &pb.CreateMovieRequest{
		Movie: data,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"movie": res.Movie,
	})
}

func (g *GinHanlder) GetByid(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := g.client.GetMovie(ctx, &pb.ReadMovieRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"movie": res.Movie,
	})
}

func (g *GinHanlder) GetAll(ctx *gin.Context) {
	res, err := g.client.GetMovies(ctx, &pb.ReadMoviesRequest{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"movies": res.Movies,
	})
}

func (g *GinHanlder) Update(ctx *gin.Context) {
	var movie models.Movies
	id := ctx.Param("id")
	err := ctx.ShouldBind(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := g.client.UpdateMovie(ctx, &pb.UpdateMovieRequest{
		Movie: &pb.Movie{
			Id:    id,
			Title: movie.Title,
			Genre: movie.Genre,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"movie": res.Movie,
	})

}

func (g *GinHanlder) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := g.client.DeleteMovie(ctx, &pb.DeleteMovieRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if res.Success == true {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Movie deleted successfully",
		})
		return
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error deleting movie",
		})
		return
	}

}
