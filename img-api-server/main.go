package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/nfnt/resize"

	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"

	fiber "github.com/gofiber/fiber/v2"
)

//go:embed images/**
var embedfile embed.FS

var SERVER_VERSION string = "v0.0.3"

type Cat struct {
	Background int16  `json:"background,omitempty"`
	Wings      int16  `json:"wings,omitempty"`
	Body       int16  `json:"body,omitempty"`
	CatType    string `json:"cat_type,omitempty"`
	Effect     int16  `json:"effect,omitempty"`
	Hat        int16  `json:"hat,omitempty"`
	Neck       int16  `json:"neck,omitempty"`
	Face       int16  `json:"face,omitempty"`
	Legs       int16  `json:"legs,omitempty"`
	Ear        int16  `json:"ear,omitempty"`
	Frame      int16  `json:"frame,omitempty"`
	Format     string `json:"format,omitempty"`
	Size       int16  `json:"size,omitempty"`
}

func init() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
}

type Item struct {
	Name string `json:"name,omitempty"`
	ID   int16  `json:"id,omitempty"`
}
type Items struct {
	Items []Item `json:"items,omitempty"`
}

func getRandTraits(itemname string) string {
	items, err := itemList(itemname)
	rand.Seed(time.Now().UnixNano())
	min := 0

	max := len([]Item(items.Items))
	log.Print(max)
	if err != nil {
		// panic(err)
		return ""
	}
	a := rand.Intn(max-min) + min
	// out:=itemname+"-"+string()
	out := fmt.Sprintf("%s-%d", itemname, a)
	return out
}
func itemList(itemname string) (Items, error) {
	data, err := embedfile.ReadFile(fmt.Sprintf("images/%s.json", itemname))
	if err != nil {
		return Items{}, err
	}
	var items Items
	err = json.Unmarshal([]byte(data), &items)
	if err != nil {
		return Items{}, err
	}
	// log.Print(items)
	return items, nil
}

func catGen(cat *Cat) (*bytes.Buffer, error) {

	format := cat.Format
	size := cat.Size

	catType := cat.CatType
	bgNumber := cat.Background
	neckNumber := cat.Neck
	hatNumber := cat.Hat

	effectNumber := cat.Effect
	bodyNumber := cat.Body
	wingsNumber := cat.Wings
	faceNumber := cat.Face
	earNumber := cat.Ear
	legsNumber := cat.Legs
	log.Printf("size=%d, format=%s, bgNumber=%d, effectNumber=%d, catType=%s, bodyNumber=%d, faceNumber=%d, wingsNumber=%d", size, format, bgNumber, effectNumber, catType, bodyNumber, faceNumber, wingsNumber)

	newImg := image.NewRGBA(image.Rect(0, 0, int(size), int(size)))
	outRect := image.Rectangle{image.Pt(0, 0), newImg.Bounds().Size()}

	/////////// Back Ground
	bgFilename := fmt.Sprintf("images/background/bg%03d.png", bgNumber)
	// log.Printf("bgFilename=%s", bgFilename)
	bg, err := embedfile.Open(bgFilename)
	if err == nil {
		defer bg.Close()

		// log.Printf("%s is exist", bgFilename)
		bgImg, _, err := image.Decode(bg)
		if err != nil {
			return nil, err
		}
		if size != 800 {
			bgImg = resize.Resize(uint(size), uint(size), bgImg, resize.NearestNeighbor)
		}
		draw.Draw(newImg, outRect, bgImg, image.Pt(0, 0), draw.Over)
	} else {
		// no background
		// return c.SendString(fmt.Sprintf("bg error %s", err))
		if format == "jpeg" {
			whitebgFilename := fmt.Sprint("images/background/white.png")
			wbg, err := embedfile.Open(whitebgFilename)
			if err == nil {
				defer wbg.Close()

				wbgImg, _, err := image.Decode(wbg)
				if err != nil {
					return nil, err
				}

				if size != 800 {
					wbgImg = resize.Resize(uint(size), uint(size), wbgImg, resize.NearestNeighbor)
				}
				draw.Draw(newImg, outRect, wbgImg, image.Pt(0, 0), draw.Over)
			}
		} else {
			// png  => nothing.

		}
	}
	/////////// Back Effect
	beFilename := fmt.Sprintf("images/back_effect/be%03d.png", effectNumber)
	// log.Printf("beFilename=%s", beFilename)

	src1, err := embedfile.Open(beFilename)
	if err == nil {
		defer src1.Close()

		src1Img, _, err := image.Decode(src1)
		if err != nil {
			return nil, err
		}
		if size != 800 {
			src1Img = resize.Resize(uint(size), uint(size), src1Img, resize.NearestNeighbor)
		}
		draw.Draw(newImg, outRect, src1Img, image.Pt(0, 0), draw.Over)
	} else {
		// no back effect
		// return c.SendString(fmt.Sprintf("back effect error %s", err))
	}
	/////////// Back Wings
	bwFilename := fmt.Sprintf("images/back_wings/bw%03d.png", wingsNumber)
	// log.Printf("beFilename=%s", beFilename)

	srcWings, err := embedfile.Open(bwFilename)
	if err == nil {
		defer srcWings.Close()

		wingsImg, _, err := image.Decode(srcWings)
		if err != nil {
			return nil, err
		}

		if size != 800 {
			wingsImg = resize.Resize(uint(size), uint(size), wingsImg, resize.NearestNeighbor)
		}
		draw.Draw(newImg, outRect, wingsImg, image.Pt(0, 0), draw.Over)
	} else {
		// no back effect
		// return c.SendString(fmt.Sprintf("back effect error %s", err))
	}
	/////////// Body
	bodyFilename := ""
	// if catType == "z" {
	// 	bodyFilename = fmt.Sprintf("images/bodies/z%03d.png", bodyNumber)
	// } else if catType == "m" {
	// 	bodyFilename = fmt.Sprintf("images/bodies/m%03d.png", bodyNumber)
	// } else if catType == "s" {
	// 	bodyFilename = fmt.Sprintf("images/bodies/s%03d.png", bodyNumber)
	// } else {
	// 	bodyFilename = fmt.Sprintf("images/bodies/cat%03d.png", bodyNumber)
	// }
	// bodyFilename = fmt.Sprintf("images/bodies/body%03d.png", bodyNumber)

	// if catType == "z" {
	// 	bodyFilename = fmt.Sprintf("images/body/z%03d.png", bodyNumber)
	// } else if catType == "m" {
	// 	bodyFilename = fmt.Sprintf("images/body/m%03d.png", bodyNumber)
	// } else if catType == "s" {
	// 	bodyFilename = fmt.Sprintf("images/body/s%03d.png", bodyNumber)
	// } else {

	// }
	bodyFilename = fmt.Sprintf("images/body/body%03d.png", bodyNumber)
	src2, err := embedfile.Open(bodyFilename)
	// log.Print(bodyFilename)
	if err == nil {
		defer src2.Close()

		src2Img, _, err := image.Decode(src2)
		if err != nil {
			return nil, err
			// return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprint(err))
		}

		if size != 800 {
			src2Img = resize.Resize(uint(size), uint(size), src2Img, resize.NearestNeighbor)
		}
		draw.Draw(newImg, outRect, src2Img, image.Pt(0, 0), draw.Over)
	} else {
		// no body

		bodyNumber = 1
		// if catType == "z" {
		// 	bodyFilename = fmt.Sprintf("images/bodies/z%03d.png", bodyNumber)
		// } else if catType == "m" {
		// 	bodyFilename = fmt.Sprintf("images/bodies/m%03d.png", bodyNumber)
		// } else if catType == "s" {
		// 	bodyFilename = fmt.Sprintf("images/bodies/s%03d.png", bodyNumber)
		// } else {
		// 	bodyFilename = fmt.Sprintf("images/bodies/cat%03d.png", bodyNumber)
		// }
		bodyFilename = fmt.Sprintf("images/body/body%03d.png", bodyNumber)

		srcBody, err := embedfile.Open(bodyFilename)
		if err == nil {
			defer srcBody.Close()

			bodyImg, _, err := image.Decode(srcBody)
			if err != nil {
				return nil, err
				// return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprint(err))
			}

			if size != 800 {
				bodyImg = resize.Resize(uint(size), uint(size), bodyImg, resize.NearestNeighbor)
			}
			draw.Draw(newImg, outRect, bodyImg, image.Pt(0, 0), draw.Over)
		} else {

			return nil, err
			// return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("no body error %s", err))
		}
	}
	/////////// Face

	faceFilename := ""
	if catType == "z" {
		faceFilename = fmt.Sprintf("images/face/z%03d.png", faceNumber)
	} else {
		faceFilename = fmt.Sprintf("images/face/f%03d.png", faceNumber)
	}
	src3, err := embedfile.Open(faceFilename)
	if err == nil {
		defer src3.Close()

		src3Img, _, err := image.Decode(src3)
		if err != nil {
			return nil, err
			// return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprint(err))
		}

		if size != 800 {
			src3Img = resize.Resize(uint(size), uint(size), src3Img, resize.NearestNeighbor)
		}
		draw.Draw(newImg, outRect, src3Img, image.Pt(0, 0), draw.Over)
	} else {
		// no faces
		// return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("faces error %s", err))

		// faceFilename = fmt.Sprint("images/face/f000.png")
		faceFilename = "images/face/f000.png"
		src3, err := embedfile.Open(faceFilename)
		if err == nil {
			defer src3.Close()

			src3Img, _, err := image.Decode(src3)
			if err != nil {
				return nil, err
				// return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("face error 1 %s", err))
			}

			if size != 800 {
				src3Img = resize.Resize(uint(size), uint(size), src3Img, resize.NearestNeighbor)
			}
			draw.Draw(newImg, outRect, src3Img, image.Pt(0, 0), draw.Over)

		} else {
			return nil, err
			// return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("face error 2 %s", err))
		}

	}
	/////////// Ear

	earFilename := ""
	earFilename = fmt.Sprintf("images/ear/ear%03d.png", earNumber)

	earSrc, err := embedfile.Open(earFilename)
	if err == nil {
		defer earSrc.Close()

		earImg, _, err := image.Decode(earSrc)
		if err != nil {
			return nil, err
			// return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprint(err))
		}

		if size != 800 {
			earImg = resize.Resize(uint(size), uint(size), earImg, resize.NearestNeighbor)
		}
		draw.Draw(newImg, outRect, earImg, image.Pt(0, 0), draw.Over)
	} else {
		// no ear
		//
	}
	/////////// Neck

	neckFilename := ""
	neckFilename = fmt.Sprintf("images/neck/neck%03d.png", neckNumber)

	neckSrc, err := embedfile.Open(neckFilename)
	if err == nil {
		defer neckSrc.Close()

		neckImg, _, err := image.Decode(neckSrc)
		if err != nil {
			return nil, err
			// return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprint(err))
		}
		if size != 800 {
			neckImg = resize.Resize(uint(size), uint(size), neckImg, resize.NearestNeighbor)
		}
		draw.Draw(newImg, outRect, neckImg, image.Pt(0, 0), draw.Over)
	} else {
		// no ear
		//
	}
	/////////// Hat

	hatFilename := ""
	hatFilename = fmt.Sprintf("images/hat/hat%03d.png", hatNumber)

	hatSrc, err := embedfile.Open(hatFilename)
	if err == nil {
		defer hatSrc.Close()

		hatImg, _, err := image.Decode(hatSrc)
		if err != nil {
			return nil, err
			// return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprint(err))
		}

		if size != 800 {
			hatImg = resize.Resize(uint(size), uint(size), hatImg, resize.NearestNeighbor)
		}
		draw.Draw(newImg, outRect, hatImg, image.Pt(0, 0), draw.Over)
	} else {
		// no ear
		//
	}
	/////////// fore Legs

	flFilename := ""
	flFilename = fmt.Sprintf("images/forelegs/forelegs%03d.png", legsNumber)

	flSrc, err := embedfile.Open(flFilename)
	if err == nil {
		defer flSrc.Close()

		flImg, _, err := image.Decode(flSrc)
		if err != nil {
			return nil, err
			// return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprint(err))
		}

		if size != 800 {
			flImg = resize.Resize(uint(size), uint(size), flImg, resize.NearestNeighbor)
		}
		draw.Draw(newImg, outRect, flImg, image.Pt(0, 0), draw.Over)
	} else {
		// no ear
		//
	}
	/////////// rear Legs

	rlFilename := ""
	rlFilename = fmt.Sprintf("images/rearlegs/rearlegs%03d.png", legsNumber)

	rlSrc, err := embedfile.Open(rlFilename)
	if err == nil {
		defer rlSrc.Close()

		rlImg, _, err := image.Decode(rlSrc)
		if err != nil {
			return nil, err
		}

		if size != 800 {
			rlImg = resize.Resize(uint(size), uint(size), rlImg, resize.NearestNeighbor)
		}
		draw.Draw(newImg, outRect, rlImg, image.Pt(0, 0), draw.Over)
	} else {
		// no ear
		//
	}
	/////////// Front Wings
	fwFilename := fmt.Sprintf("images/front_wings/fw%03d.png", wingsNumber)
	// log.Printf("beFilename=%s", beFilename)

	srcFrontWings, err := embedfile.Open(fwFilename)
	if err == nil {
		defer srcFrontWings.Close()

		frontWingsImg, _, err := image.Decode(srcFrontWings)
		if err != nil {
			return nil, err
		}
		if size != 800 {
			frontWingsImg = resize.Resize(uint(size), uint(size), frontWingsImg, resize.NearestNeighbor)
		}
		draw.Draw(newImg, outRect, frontWingsImg, image.Pt(0, 0), draw.Over)
	} else {
		// no back effect
		// return c.SendString(fmt.Sprintf("back effect error %s", err))
	}
	/////////// Front Effect

	feFilename := fmt.Sprintf("images/front_effect/fe%03d.png", effectNumber)
	src4, err := embedfile.Open(feFilename)
	if err == nil {
		defer src4.Close()

		src4Img, _, err := image.Decode(src4)
		if err != nil {
			return nil, err

			// return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprint(err))
		}
		if size != 800 {
			src4Img = resize.Resize(uint(size), uint(size), src4Img, resize.NearestNeighbor)
		}
		draw.Draw(newImg, outRect, src4Img, image.Pt(0, 0), draw.Over)
	} else {
		// no front effect
		// return c.SendString(fmt.Sprintf("Front Effect error %s", err))

	}

	/////////// Frame

	frameFilename := fmt.Sprintf("images/frame/frame%03d.png", cat.Frame)
	// log.Printf("beFilename=%s", beFilename)

	srcFrame, err := embedfile.Open(frameFilename)
	if err == nil {
		defer srcFrame.Close()

		frameImg, _, err := image.Decode(srcFrame)
		if err != nil {
			return nil, err
		}
		if size != 800 {
			frameImg = resize.Resize(uint(size), uint(size), frameImg, resize.NearestNeighbor)
		}
		draw.Draw(newImg, outRect, frameImg, image.Pt(0, 0), draw.Over)
	} else {
		// no frame
		// return c.SendString(fmt.Sprintf("back effect error %s", err))
	}
	// output := "jpeg" // png or jpeg
	if format == "jpeg" {
		//jpeg
		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, newImg, nil); err != nil {
			return nil, err
			// return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprint("jpeg.Encode Error. ( %s )", err))
		}
		return buffer, nil

		// msg := fmt.Sprint(buffer)
		// c.Set("Content-type", "image/jpeg")
		// return c.SendString(msg)
	}
	// png
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, newImg); err != nil {
		return nil, err
		// return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprint("png.Encode Error. ( %s )", err))
	}
	// msg := fmt.Sprint(buffer)
	// c.Set("Content-type", "image/png")
	// return c.SendString(msg)

	return buffer, nil
}
func main() {

	//ref: https://github.com/gofiber/template
	// engine := html.New("./views", ".html")
	// engine := html.NewFileSystem(http.Dir("./views"), ".html")
	// engine := html.NewFileSystem(rice.MustFindBox("views").HTTPBox(), ".html")

	// Reload the templates on each render, good for development
	// engine.Reload(true)    // Optional. Default: false
	// engine.Layout("embed") // Optional. Default: "embed"

	// app := fiber.New()
	app := fiber.New(fiber.Config{
		// Prefork: true,
		// CaseSensitive: true,
		// StrictRouting: true,
		ServerHeader: "Fiber",
		// Views:        engine,
		AppName: fmt.Sprintf("Twinkle Image App Server %s", SERVER_VERSION),
	})
	// https://docs.gofiber.io/api/middleware/recover
	app.Use(recover.New())

	// compress
	app.Use(compress.New(compress.Config{
		// Skip compress  for specific routes
		// Next: func(c *fiber.Ctx) bool {
		// 	return c.Path() == "/v0/img" // no compress
		// },
		Level: compress.LevelBestSpeed, // 1
	}))
	// etag.
	app.Use(etag.New())
	// cache. ref: https://docs.gofiber.io/api/middleware/cache

	// logging. ref: https://docs.gofiber.io/api/middleware/logger
	// app.Use(cache.New())
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return strings.HasPrefix(c.Path(), "/v0/img/randomcat")
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Get("/", func(c *fiber.Ctx) error {
		// msg := "hello human🐱 "
		data, err := embedfile.ReadFile("images/index.html")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error(frame)  %s", err))
		}
		c.Set("Content-type", "text/html; charset=utf8")
		// return c.SendString(msg + string(data))
		return c.SendString(string(data))
	})
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Render("index", fiber.Map{
	// 		// "Title": "Hello, World!",
	// 		"message": "hello human🐱 ",
	// 		"version": SERVER_VERSION,
	// 	}, "layouts/main")
	// })

	// GET /
	app.Get("/v0/api/items",
		func(c *fiber.Ctx) error {
			out := make(map[string]Items)
			hat, err := itemList("hat")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error(hat)  %s", err))
			}
			out["hat"] = hat

			frame, err := itemList("frame")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error(frame)  %s", err))
			}
			out["frame"] = frame

			data, err := itemList("neck")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error(neck)  %s", err))
			}
			out["neck"] = data

			out["hat"] = hat

			body, err := itemList("body")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error(body)  %s", err))
			}
			out["body"] = body

			legs, err := itemList("legs")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error(legs)  %s", err))
			}
			out["legs"] = legs

			wings, err := itemList("wings")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error(wings)  %s", err))
			}
			out["wings"] = wings

			effect, err := itemList("effect")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error(effect)  %s", err))
			}
			out["effect"] = effect

			face, err := itemList("face")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error(face)  %s", err))
			}
			out["face"] = face

			ear, err := itemList("ear")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error(ear)  %s", err))
			}
			out["ear"] = ear

			bg, err := itemList("bg")
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error(bg)  %s", err))
			}
			out["bg"] = bg

			jout, err := json.Marshal(out)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error(marshall)  %s", err))
			}

			return c.SendString(string(jout))
		},
	)

	// /v0/img/cat/png/bg-12.body-4.wings-1.hat-1.legs-1.effect-1.frame-1.type-m.neck-1.face-1.ear-1
	app.Get("/v0/img/randomcat/*", func(c *fiber.Ctx) error {
		args := strings.Split(c.Params("*"), ".")
		wings := int16(-1)

		body := int16(-1) //
		bg := int16(-1)   //
		effect := int16(-1)
		hat := int16(-1)
		neck := int16(-1)
		legs := int16(-1)
		// frame := int16(-1)
		face := int16(-1)
		ear := int16(-1)
		for _, arg := range args {
			// log.Print(arg)

			namevalue := strings.Split(arg, "-")
			switch namevalue[0] {
			case "face": // face
				num, err := strconv.ParseInt(namevalue[1], 10, 0)

				if err == nil {
					face = int16(num)
				} else {
					face = -1 // random
				}
			case "ear": // ear
				num, err := strconv.ParseInt(namevalue[1], 10, 0)

				if err == nil {
					ear = int16(num)
				} else {
					ear = -1 // random
				}
			case "wings": // wings
				num, err := strconv.ParseInt(namevalue[1], 10, 0)

				if err == nil {
					wings = int16(num)
				} else {
					wings = -1 // random
				}
			case "body": // body
				num, err := strconv.ParseInt(namevalue[1], 10, 0)

				if err == nil {
					body = int16(num)
				} else {
					body = -1 // random
				}
			case "bg": // bg
				num, err := strconv.ParseInt(namevalue[1], 10, 0)

				if err == nil {
					bg = int16(num)
				} else {
					bg = -1 // random
				}
			case "effect": // effect
				num, err := strconv.ParseInt(namevalue[1], 10, 0)

				if err == nil {
					effect = int16(num)
				} else {
					effect = -1 // random
				}
			case "hat": // hat
				num, err := strconv.ParseInt(namevalue[1], 10, 0)

				if err == nil {
					hat = int16(num)
				} else {
					hat = -1 // random
				}
			case "neck": // neck
				num, err := strconv.ParseInt(namevalue[1], 10, 0)

				if err == nil {
					neck = int16(num)
				} else {
					neck = -1 // random
				}
			case "legs": // legs
				num, err := strconv.ParseInt(namevalue[1], 10, 0)

				if err == nil {
					legs = int16(num)
				} else {
					legs = -1 // random
				}
			}
		}
		traitItems := []string{}
		if wings < 0 {
			traitItems = append(traitItems, getRandTraits("wings"))
		} else {
			traitItems = append(traitItems, fmt.Sprintf("wings-%d", wings))
		}
		if bg < 0 {
			traitItems = append(traitItems, getRandTraits("bg"))
		} else {
			traitItems = append(traitItems, fmt.Sprintf("bg-%d", bg))
		}
		if body < 0 {
			traitItems = append(traitItems, getRandTraits("body"))
		} else {
			traitItems = append(traitItems, fmt.Sprintf("body-%d", body))
		}
		if effect < 0 {
			traitItems = append(traitItems, getRandTraits("effect"))
		} else {
			traitItems = append(traitItems, fmt.Sprintf("effect-%d", effect))
		}
		if hat < 0 {
			traitItems = append(traitItems, getRandTraits("hat"))
		} else {
			traitItems = append(traitItems, fmt.Sprintf("hat-%d", hat))
		}

		if ear < 0 {
			traitItems = append(traitItems, getRandTraits("ear"))
		} else {
			traitItems = append(traitItems, fmt.Sprintf("ear-%d", ear))
		}

		if legs < 0 {
			traitItems = append(traitItems, getRandTraits("legs"))
		} else {
			traitItems = append(traitItems, fmt.Sprintf("legs-%d", legs))
		}

		if neck < 0 {
			traitItems = append(traitItems, getRandTraits("neck"))
		} else {
			traitItems = append(traitItems, fmt.Sprintf("neck-%d", neck))
		}

		if face < 0 {
			traitItems = append(traitItems, getRandTraits("face"))
		} else {
			traitItems = append(traitItems, fmt.Sprintf("face-%d", face))
		}
		traitItems = append(traitItems, "size-300")

		// traitItems = append(traitItems, getRandTraits("body"))
		log.Print(traitItems)
		c.Set("Content-type", "text/html; charset=utf8")
		a := strings.Join(traitItems, ".")
		// return c.SendString(fmt.Sprintf("rand  %v", traitItems))
		imgurl := fmt.Sprintf("%s://%s/v0/img/cat/jpg/%s", c.Protocol(), c.Hostname(), a)
		return c.Redirect(imgurl)

	})

	// /v0/img/cat/png/bg-12.body-4.wings-1.hat-1.legs-1.effect-1.frame-1.type-m.neck-1.face-1.ear-1
	app.Get("/v0/img/cat/:format/*", func(c *fiber.Ctx) error {
		format := c.Params("format")
		if format != "png" {
			format = "jpeg"
		}
		catType := "c"
		wings := int16(0) //
		body := int16(1)  //
		bg := int16(0)    //
		effect := int16(0)
		hat := int16(0)
		neck := int16(0)
		legs := int16(0)
		frame := int16(0)
		face := int16(0)
		ear := int16(0)
		size := int16(800)
		// c.SendString(c.Params("+"))
		args := strings.Split(c.Params("*"), ".")
		for _, arg := range args {
			// log.Print(arg)
			namevalue := strings.Split(arg, "-")
			switch namevalue[0] {
			case "size": // size
				num, err := strconv.ParseInt(namevalue[1], 10, 0)
				if err == nil {
					if num == 300 || num == 400 || num == 500 || num == 600 {
						size = int16(num)
					}
				} else {
					size = 800
				}

			case "wings": // wings
				num, err := strconv.ParseInt(namevalue[1], 10, 0)
				if err == nil {
					wings = int16(num)
				} else {
					wings = -1 // random
				}

			case "ear": // ear
				num, err := strconv.ParseInt(namevalue[1], 10, 0)
				if err == nil {
					ear = int16(num)
				} else {
					ear = -1 // random
				}
			case "face": // face
				num, err := strconv.ParseInt(namevalue[1], 10, 0)
				if err == nil {
					face = int16(num)
				} else {
					face = -1 // random
				}

			case "bg": // background
				num, err := strconv.ParseInt(namevalue[1], 10, 0)
				if err == nil {
					bg = int16(num)
				} else {
					bg = -1 // random
				}

			case "body": // body
				num, err := strconv.ParseInt(namevalue[1], 10, 0)
				if err == nil {
					body = int16(num)
				} else {
					body = -1 // random
				}

			case "effect": // effect
				num, err := strconv.ParseInt(namevalue[1], 10, 0)
				if err == nil {
					effect = int16(num)
				} else {
					effect = -1 // random
				}
			case "frame": // frame
				num, err := strconv.ParseInt(namevalue[1], 10, 0)
				if err == nil {
					frame = int16(num)
				} else {
					frame = -1 // random
				}
			case "legs": // legs
				num, err := strconv.ParseInt(namevalue[1], 10, 0)
				if err == nil {
					legs = int16(num)
				} else {
					legs = -1 // random
				}
			case "neck": // neck
				num, err := strconv.ParseInt(namevalue[1], 10, 0)
				if err == nil {
					neck = int16(num)
				} else {
					neck = -1 // random
				}
			case "hat": // hat
				num, err := strconv.ParseInt(namevalue[1], 10, 0)
				if err == nil {
					hat = int16(num)
				} else {
					hat = -1 // random
				}
				// case "type": // catType
				// 	switch namevalue[1][0] {
				// 	case 'm':
				// 		catType = "m"
				// 	case 'z':
				// 		catType = "z"
				// 	case 'c':
				// 		catType = "c"
				// 	case 's':
				// 		catType = "s"
				// 	}
			}
		}
		cat := Cat{Size: size, Format: format, CatType: catType, Wings: wings, Background: bg, Body: body, Frame: frame, Neck: neck, Hat: hat, Legs: legs, Effect: effect, Face: face, Ear: ear}

		buffer, err := catGen(&cat)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error  %s", err))
		}

		// output := "jpeg" // png or jpeg
		if format == "jpeg" {
			msg := fmt.Sprint(buffer)
			c.Set("Content-type", "image/jpeg")
			return c.SendString(msg)
		}
		// png

		c.Set("Content-type", "image/png")
		// c.Set("Content-type", "image/png")
		return c.SendString(fmt.Sprint(buffer))
	})
	// GET /img/cat/png/19/4/c/3/1/1
	// GET /img/cat/png/5/6/s/6/1/1
	// GET /img/cat/png/1/10/m/1/1/1
	// GET /img/cat/png/5/19/z/1/1/1
	// img/cat/png/5/4/z/1/3/2
	app.Get("/img/cat/:format/:bgStr?/:effectStr?/:bodyType?/:bodyStr?/:faceStr?/:wingsStr?", func(c *fiber.Ctx) error {
		format := c.Params("format")
		if format != "png" {
			format = "jpeg"
		}

		bodyType := c.Params("bodyType")

		bgNumber, err := strconv.Atoi(c.Params("bgStr"))
		if err != nil {
			log.Print(err)
		}
		effectNumber, err := strconv.Atoi(c.Params("effectStr"))
		if err != nil {
			log.Print(err)
		}
		bodyNumber, err := strconv.Atoi(c.Params("bodyStr"))
		if err != nil {
			log.Print(err)
		}
		faceNumber, err := strconv.Atoi(c.Params("faceStr"))
		if err != nil {
			log.Print(err)
		}
		wingsNumber, err := strconv.Atoi(c.Params("wingsStr"))
		if err != nil {
			log.Print(err)
		}
		log.Printf("format=%s, bgNumber=%d, effectNumber=%d, bodyType=%s, bodyNumber=%d, faceNumber=%d, wingsNumber=%d", format, bgNumber, effectNumber, bodyType, bodyNumber, faceNumber, wingsNumber)

		cat := Cat{Format: format, CatType: bodyType, Wings: int16(wingsNumber), Background: int16(bgNumber), Body: int16(bodyNumber), Frame: 0, Neck: 0, Hat: 0, Legs: 0, Effect: int16(effectNumber), Face: int16(faceNumber), Ear: 0}

		buffer, err := catGen(&cat)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error  %s", err))
		}

		// output := "jpeg" // png or jpeg
		if format == "jpeg" {
			msg := fmt.Sprint(buffer)
			c.Set("Content-type", "image/jpeg")
			return c.SendString(msg)
		}
		// png

		c.Set("Content-type", "image/png")
		// c.Set("Content-type", "image/png")
		return c.SendString(fmt.Sprint(buffer))

	})

	log.Fatal(app.Listen(":3000"))
}
