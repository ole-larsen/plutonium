package grpcserver

import (
	frontendv1 "github.com/ole-larsen/plutonium/gen/frontend/v1"
	marketv1 "github.com/ole-larsen/plutonium/gen/market/v1"
	"github.com/ole-larsen/plutonium/models"
)

func NestPublicMenu(items []*models.PublicMenu) []*frontendv1.PublicMenu {
	nestedItems := make([]*frontendv1.PublicMenu, len(items))

	for i := 0; i < len(items); i++ {
		nestedItems[i] = &frontendv1.PublicMenu{
			Id: items[i].ID,
			Attributes: &frontendv1.PublicMenuAttributes{
				Name:    items[i].Attributes.Name,
				Link:    items[i].Attributes.Link,
				OrderBy: items[i].Attributes.OrderBy,
				Items:   NestPublicMenu(items[i].Attributes.Items),
			},
		}
	}

	return nestedItems
}

func NestFrontendPublicFile(file *models.PublicFile) *frontendv1.PublicFile {
	if file == nil {
		return nil
	}

	return &frontendv1.PublicFile{
		Id: file.ID,
		Attributes: &frontendv1.PublicFileAttributes{
			Name:     file.Attributes.Name,
			Provider: file.Attributes.Provider,
			Url:      file.Attributes.URL,
			Alt:      file.Attributes.Alt,
			Caption:  file.Attributes.Caption,
			Ext:      file.Attributes.Ext,
			Hash:     file.Attributes.Hash,
			Mime:     file.Attributes.Mime,
			Size:     file.Attributes.Size,
			Width:    file.Attributes.Width,
			Height:   file.Attributes.Height,
		},
	}
}

func NestMarketPublicFile(file *models.PublicFile) *marketv1.PublicFile {
	if file == nil {
		return nil
	}

	return &marketv1.PublicFile{
		Id: file.ID,
		Attributes: &marketv1.PublicFileAttributes{
			Name:     file.Attributes.Name,
			Provider: file.Attributes.Provider,
			Url:      file.Attributes.URL,
			Alt:      file.Attributes.Alt,
			Caption:  file.Attributes.Caption,
			Ext:      file.Attributes.Ext,
			Hash:     file.Attributes.Hash,
			Mime:     file.Attributes.Mime,
			Size:     file.Attributes.Size,
			Width:    file.Attributes.Width,
			Height:   file.Attributes.Height,
		},
	}
}

func NestFrontendPublicUser(user *models.PublicUser) *frontendv1.PublicUser {
	nestedUser := &frontendv1.PublicUser{
		Id: user.ID,
		Attributes: &frontendv1.PublicUserAttributes{
			Username: user.Attributes.Username,
			Address:  user.Attributes.Address,
			Email:    user.Attributes.Email,
			Uuid:     user.Attributes.UUID,
			Gravatar: user.Attributes.Gravatar,
			Nonce:    user.Attributes.Nonce,
			Token:    user.Attributes.Token,
			Funds:    user.Attributes.Funds,
			Created:  user.Attributes.Created,
		},
	}
	if user.Attributes.Wallpaper != nil {
		nestedUser.Attributes.Wallpaper = NestFrontendPublicFile(user.Attributes.Wallpaper)
	}

	return nestedUser
}

func NestMarketPublicUser(user *models.PublicUser) *marketv1.PublicUser {
	nestedUser := &marketv1.PublicUser{
		Id: user.ID,
		Attributes: &marketv1.PublicUserAttributes{
			Username: user.Attributes.Username,
			Address:  user.Attributes.Address,
			Email:    user.Attributes.Email,
			Uuid:     user.Attributes.UUID,
			Gravatar: user.Attributes.Gravatar,
			Nonce:    user.Attributes.Nonce,
			Token:    user.Attributes.Token,
			Funds:    user.Attributes.Funds,
			Created:  user.Attributes.Created,
		},
	}
	if user.Attributes.Wallpaper != nil {
		nestedUser.Attributes.Wallpaper = NestMarketPublicFile(user.Attributes.Wallpaper)
	}

	return nestedUser
}

func NestPublicSlides(items []*models.PublicSliderItem) []*frontendv1.PublicSliderItem {
	nestedItems := make([]*frontendv1.PublicSliderItem, len(items))

	for i := 0; i < len(items); i++ {
		nestedItems[i] = &frontendv1.PublicSliderItem{
			Id:          items[i].ID,
			Description: items[i].Description,
			Heading:     items[i].Heading,
			BtnLink1:    items[i].BtnLink1,
			BtnLink2:    items[i].BtnLink2,
			BtnText1:    items[i].BtnText1,
			BtnText2:    items[i].BtnText2,
		}
		if items[i].Image != nil {
			nestedItems[i].Image = NestFrontendPublicFile(items[i].Image)
		}

		if items[i].Bg != nil {
			nestedItems[i].Bg = NestFrontendPublicFile(items[i].Bg)
		}
	}

	return nestedItems
}

func NestPublicCategories(items []*models.PublicCategory) []*marketv1.PublicCategory {
	nestedItems := make([]*marketv1.PublicCategory, len(items))
	for i := 0; i < len(items); i++ {
		nestedItems[i] = &marketv1.PublicCategory{
			Id: items[i].ID,
			Attributes: &marketv1.PublicCategoryAttributes{
				Title:       items[i].Attributes.Title,
				Slug:        items[i].Attributes.Slug,
				Description: items[i].Attributes.Description,
				Content:     items[i].Attributes.Content,
			},
		}
		if items[i].Attributes.Image != nil {
			nestedItems[i].Attributes.Image = NestMarketPublicFile(items[i].Attributes.Image)
		}

		if items[i].Attributes.Collections != nil {
			nestedItems[i].Attributes.Collections = NestMarketCollections(items[i].Attributes.Collections)
		}
	}

	return nestedItems
}

func NestPublicCreateAndSellItems(items []*models.PublicCreateAndSellItem) []*frontendv1.PublicCreateAndSellItem {
	nestedItems := make([]*frontendv1.PublicCreateAndSellItem, len(items))
	for i := 0; i < len(items); i++ {
		nestedItems[i] = &frontendv1.PublicCreateAndSellItem{
			Id: items[i].ID,
			Attributes: &frontendv1.PublicCreateAndSellItemAttributes{
				Title:       items[i].Attributes.Title,
				Link:        items[i].Attributes.Link,
				Description: items[i].Attributes.Description,
				Image:       NestFrontendPublicFile(items[i].Attributes.Image),
			},
		}
		if items[i].Attributes.Image != nil {
			nestedItems[i].Attributes.Image = NestFrontendPublicFile(items[i].Attributes.Image)
		}
	}

	return nestedItems
}

func NestPublicHelpCenterItems(items []*models.PublicHelpCenterItem) []*frontendv1.PublicHelpCenterItem {
	nestedItems := make([]*frontendv1.PublicHelpCenterItem, len(items))
	for i := 0; i < len(items); i++ {
		nestedItems[i] = &frontendv1.PublicHelpCenterItem{
			Id: items[i].ID,
			Attributes: &frontendv1.PublicHelpCenterItemAttributes{
				Title:       items[i].Attributes.Title,
				Link:        items[i].Attributes.Link,
				Description: items[i].Attributes.Description,
				Image:       NestFrontendPublicFile(items[i].Attributes.Image),
			},
		}
		if items[i].Attributes.Image != nil {
			nestedItems[i].Attributes.Image = NestFrontendPublicFile(items[i].Attributes.Image)
		}
	}

	return nestedItems
}

func NestPublicFaqItems(items []*models.PublicFaqItem) []*frontendv1.PublicFaqItem {
	nestedItems := make([]*frontendv1.PublicFaqItem, len(items))
	for i := 0; i < len(items); i++ {
		nestedItems[i] = &frontendv1.PublicFaqItem{
			Id: items[i].ID,
			Attributes: &frontendv1.PublicFaqItemAttributes{
				Answer:   items[i].Attributes.Answer,
				Question: items[i].Attributes.Question,
			},
		}
	}

	return nestedItems
}

func NestPublicContact(item *models.PublicContact) *frontendv1.PublicContact {
	if item == nil {
		return nil
	}

	nestedItem := &frontendv1.PublicContact{
		Id: item.ID,
		Attributes: &frontendv1.PublicContactAttributes{
			Heading:    item.Attributes.Heading,
			SubHeading: item.Attributes.SubHeading,
			Link:       item.Attributes.Link,
			Text:       item.Attributes.Text,
			Csrf:       item.Attributes.Csrf,
		},
	}
	if item.Attributes.Image != nil {
		nestedItem.Attributes.Image = NestFrontendPublicFile(item.Attributes.Image)
	}

	return nestedItem
}

func NestPublicPage(item *models.PublicPage) *frontendv1.PublicPage {
	if item == nil {
		return nil
	}

	nestedItem := &frontendv1.PublicPage{
		Id: item.ID,
		Attributes: &frontendv1.PublicPageAttributes{
			Title:       item.Attributes.Title,
			Content:     item.Attributes.Content,
			Description: item.Attributes.Description,
			Link:        item.Attributes.Link,
			Category:    item.Attributes.Category,
		},
	}
	if item.Attributes.Image != nil {
		nestedItem.Attributes.Image = NestFrontendPublicFile(item.Attributes.Image)
	}

	return nestedItem
}

func NestPublicWalletConnectItems(items []*models.PublicWalletConnectItem) []*frontendv1.PublicWalletConnectItem {
	nestedItems := make([]*frontendv1.PublicWalletConnectItem, len(items))
	for i := 0; i < len(items); i++ {
		nestedItems[i] = &frontendv1.PublicWalletConnectItem{
			Id: items[i].ID,
			Attributes: &frontendv1.PublicWalletConnectItemAttributes{
				Title:       items[i].Attributes.Title,
				Link:        items[i].Attributes.Link,
				Description: items[i].Attributes.Description,
				Image:       NestFrontendPublicFile(items[i].Attributes.Image),
			},
		}
		if items[i].Attributes.Image != nil {
			nestedItems[i].Attributes.Image = NestFrontendPublicFile(items[i].Attributes.Image)
		}
	}

	return nestedItems
}

func NestMetadataAttributes(items []*models.MetadataAttributes) []*marketv1.MetadataAttributes {
	nestedItems := make([]*marketv1.MetadataAttributes, len(items))
	for i := 0; i < len(items); i++ {
		nestedItems[i] = &marketv1.MetadataAttributes{
			TraitType: items[i].TraitType,
			Value:     items[i].Value,
		}
	}

	return nestedItems
}

func NestMarketCollectibles(items []*models.MarketplaceCollectible) []*marketv1.MarketplaceCollectible {
	nestedItems := make([]*marketv1.MarketplaceCollectible, len(items))
	for i := 0; i < len(items); i++ {
		nestedItems[i] = &marketv1.MarketplaceCollectible{
			Id: items[i].ID,
			Attributes: &marketv1.MarketplaceCollectibleAttributes{
				CollectionId: items[i].Attributes.CollectionID,
				TokenIds:     items[i].Attributes.TokenIds,
				Uri:          items[i].Attributes.URI,
				Details: &marketv1.MarketplaceCollectibleDetails{
					Address:         items[i].Attributes.Details.Address,
					Auction:         items[i].Attributes.Details.Auction,
					Cancelled:       items[i].Attributes.Details.Cancelled,
					Collection:      items[i].Attributes.Details.Collection,
					EndTime:         items[i].Attributes.Details.EndTime,
					Fee:             items[i].Attributes.Details.Fee,
					FeeWei:          items[i].Attributes.Details.FeeWei,
					Fulfilled:       items[i].Attributes.Details.Fulfilled,
					IsStarted:       items[i].Attributes.Details.IsStarted,
					Price:           items[i].Attributes.Details.Price,
					PriceWei:        items[i].Attributes.Details.PriceWei,
					Quantity:        items[i].Attributes.Details.Quantity,
					ReservePrice:    items[i].Attributes.Details.ReservePrice,
					ReservePriceWei: items[i].Attributes.Details.ReservePriceWei,
					StartPrice:      items[i].Attributes.Details.StartPrice,
					StartPriceWei:   items[i].Attributes.Details.StartPriceWei,
					StartTime:       items[i].Attributes.Details.StartTime,
					Tags:            items[i].Attributes.Details.Tags,
					Total:           items[i].Attributes.Details.Total,
					TotalWei:        items[i].Attributes.Details.TotalWei,
				},
				Metadata: &marketv1.MarketplaceCollectibleMetadata{
					ExternalUrl:     items[i].Attributes.Metadata.ExternalURL,
					AnimationUrl:    items[i].Attributes.Metadata.AnimationURL,
					BackgroundColor: items[i].Attributes.Metadata.BackgroundColor,
					Description:     items[i].Attributes.Metadata.Description,
					YoutubeUrl:      items[i].Attributes.Metadata.YoutubeURL,
					ImageUrl:        items[i].Attributes.Metadata.ImageURL,
					Attributes:      NestMetadataAttributes(items[i].Attributes.Metadata.Attributes),
				},
			},
		}
		if items[i].Attributes.Creator != nil {
			nestedItems[i].Attributes.Creator = NestMarketPublicUser(items[i].Attributes.Creator)
		}

		if items[i].Attributes.Owner != nil {
			nestedItems[i].Attributes.Owner = NestMarketPublicUser(items[i].Attributes.Owner)
		}
	}

	return nestedItems
}

func NestMarketCollections(items []*models.MarketplaceCollection) []*marketv1.MarketplaceCollection {
	nestedItems := make([]*marketv1.MarketplaceCollection, len(items))
	for i := 0; i < len(items); i++ {
		nestedItems[i] = &marketv1.MarketplaceCollection{
			Id: items[i].ID,
			Attributes: &marketv1.MarketplaceCollectionAttributes{
				CategoryId:  items[i].Attributes.CategoryID,
				Name:        items[i].Attributes.Name,
				Slug:        items[i].Attributes.Slug,
				Url:         items[i].Attributes.URL,
				Symbol:      items[i].Attributes.Symbol,
				Description: items[i].Attributes.Description,
				Fee:         items[i].Attributes.Fee,
				Address:     items[i].Attributes.Address.String(),
				MaxItems:    items[i].Attributes.MaxItems,
				IsApproved:  items[i].Attributes.IsApproved,
				IsLocked:    items[i].Attributes.IsLocked,
				Created:     items[i].Attributes.Created,
			},
		}
		if items[i].Attributes.Logo != nil {
			nestedItems[i].Attributes.Logo = NestMarketPublicFile(items[i].Attributes.Logo)
		}

		if items[i].Attributes.Banner != nil {
			nestedItems[i].Attributes.Banner = NestMarketPublicFile(items[i].Attributes.Banner)
		}

		if items[i].Attributes.Featured != nil {
			nestedItems[i].Attributes.Featured = NestMarketPublicFile(items[i].Attributes.Featured)
		}

		if items[i].Attributes.Creator != nil {
			nestedItems[i].Attributes.Creator = NestMarketPublicUser(items[i].Attributes.Creator)
		}

		if items[i].Attributes.Owner != nil {
			nestedItems[i].Attributes.Creator = NestMarketPublicUser(items[i].Attributes.Owner)
		}

		if items[i].Attributes.Collectibles != nil {
			nestedItems[i].Attributes.Collectibles = NestMarketCollectibles(items[i].Attributes.Collectibles)
		}
	}

	return nestedItems
}
