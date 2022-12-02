package notion

import (
	"encoding/json"
	"time"
)

// Block represents content on the Notion platform.
// See: https://developers.notion.com/reference/block
type Block interface {
	ID() string
	Parent() Parent
	CreatedTime() time.Time
	CreatedBy() BlockUser
	LastEditedBy() BlockUser
	LastEditedTime() time.Time
	HasChildren() bool
	Archived() bool
}

type BlockDTO struct {
	BaseBlock
	Type BlockType `json:"type,omitempty"`

	Paragraph        *ParagraphBlock        `json:"paragraph,omitempty"`
	Heading1         *Heading1Block         `json:"heading_1,omitempty"`
	Heading2         *Heading2Block         `json:"heading_2,omitempty"`
	Heading3         *Heading3Block         `json:"heading_3,omitempty"`
	BulletedListItem *BulletedListItemBlock `json:"bulleted_list_item,omitempty"`
	NumberedListItem *NumberedListItemBlock `json:"numbered_list_item,omitempty"`
	ToDo             *ToDoBlock             `json:"to_do,omitempty"`
	Toggle           *ToggleBlock           `json:"toggle,omitempty"`
	ChildPage        *ChildPageBlock        `json:"child_page,omitempty"`
	ChildDatabase    *ChildDatabaseBlock    `json:"child_database,omitempty"`
	Callout          *CalloutBlock          `json:"callout,omitempty"`
	Quote            *QuoteBlock            `json:"quote,omitempty"`
	Code             *CodeBlock             `json:"code,omitempty"`
	Embed            *EmbedBlock            `json:"embed,omitempty"`
	Image            *ImageBlock            `json:"image,omitempty"`
	Audio            *AudioBlock            `json:"audio,omitempty"`
	Video            *VideoBlock            `json:"video,omitempty"`
	File             *FileBlock             `json:"file,omitempty"`
	PDF              *PDFBlock              `json:"pdf,omitempty"`
	Bookmark         *BookmarkBlock         `json:"bookmark,omitempty"`
	Equation         *EquationBlock         `json:"equation,omitempty"`
	Divider          *DividerBlock          `json:"divider,omitempty"`
	TableOfContents  *TableOfContentsBlock  `json:"table_of_contents,omitempty"`
	Breadcrumb       *BreadcrumbBlock       `json:"breadcrumb,omitempty"`
	ColumnList       *ColumnListBlock       `json:"column_list,omitempty"`
	Column           *ColumnBlock           `json:"column,omitempty"`
	Table            *TableBlock            `json:"table,omitempty"`
	TableRow         *TableRowBlock         `json:"table_row,omitempty"`
	LinkPreview      *LinkPreviewBlock      `json:"link_preview,omitempty"`
	LinkToPage       *LinkToPageBlock       `json:"link_to_page,omitempty"`
	SyncedBlock      *SyncedBlock           `json:"synced_block,omitempty"`
	Template         *TemplateBlock         `json:"template,omitempty"`
}

// func (b BlockDTO) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(&b)
// }

type BlockUser struct {
	Object string `json:"object"`
	ID     string `json:"id"`
}

type BaseBlock struct {
	BID             string    `json:"id,omitempty"`
	BParent         Parent    `json:"parent,omitempty"`
	BCreatedTime    time.Time `json:"created_time,omitempty"`
	BCreatedBy      BlockUser `json:"created_by,omitempty"`
	BLastEditedTime time.Time `json:"last_edited_time,omitempty"`
	BLastEditedBy   BlockUser `json:"last_edited_by,omitempty"`
	BHasChildren    bool      `json:"has_children,omitempty"`
	BArchived       bool      `json:"archived,omitempty"`
}

// ID returns the identifier (UUIDv4) for the block.
func (b BaseBlock) ID() string {
	return b.BID
}

func (b BaseBlock) CreatedTime() time.Time {
	return b.BCreatedTime
}

func (b BaseBlock) CreatedBy() BlockUser {
	return b.BCreatedBy
}

func (b BaseBlock) LastEditedTime() time.Time {
	return b.BLastEditedTime
}

func (b BaseBlock) LastEditedBy() BlockUser {
	return b.BLastEditedBy
}

func (b BaseBlock) HasChildren() bool {
	return b.BHasChildren
}

func (b BaseBlock) Archived() bool {
	return b.BArchived
}

func (b BaseBlock) Parent() Parent {
	return b.BParent
}

type ParagraphBlock struct {
	RichText []RichText `json:"rich_text"`
	Children []Block    `json:"children,omitempty"`
	Color    Color      `json:"color,omitempty"`
}

type BulletedListItemBlock struct {
	RichText []RichText `json:"rich_text"`
	Children []Block    `json:"children,omitempty"`
	Color    Color      `json:"color,omitempty"`
}

type NumberedListItemBlock struct {
	RichText []RichText `json:"rich_text"`
	Children []Block    `json:"children,omitempty"`
	Color    Color      `json:"color,omitempty"`
}

type QuoteBlock struct {
	RichText []RichText `json:"rich_text"`
	Children []Block    `json:"children,omitempty"`
	Color    Color      `json:"color,omitempty"`
}

type ToggleBlock struct {
	RichText []RichText `json:"rich_text"`
	Children []Block    `json:"children,omitempty"`
	Color    Color      `json:"color,omitempty"`
}

type TemplateBlock struct {
	RichText []RichText `json:"rich_text"`
	Children []Block    `json:"children,omitempty"`
}

type Heading1Block struct {
	RichText     []RichText `json:"rich_text"`
	Children     []Block    `json:"children,omitempty"`
	Color        Color      `json:"color,omitempty"`
	IsToggleable bool       `json:"is_toggleable"`
}

type Heading2Block struct {
	RichText     []RichText `json:"rich_text"`
	Children     []Block    `json:"children,omitempty"`
	Color        Color      `json:"color,omitempty"`
	IsToggleable bool       `json:"is_toggleable"`
}

type Heading3Block struct {
	RichText     []RichText `json:"rich_text"`
	Children     []Block    `json:"children,omitempty"`
	Color        Color      `json:"color,omitempty"`
	IsToggleable bool       `json:"is_toggleable"`
}

type ToDoBlock struct {
	RichText []RichText `json:"rich_text"`
	Children []Block    `json:"children,omitempty"`
	Checked  *bool      `json:"checked,omitempty"`
	Color    Color      `json:"color,omitempty"`
}

type ChildPageBlock struct {
	Title string `json:"title"`
}

type ChildDatabaseBlock struct {
	Title string `json:"title"`
}

type CalloutBlock struct {
	RichText []RichText `json:"rich_text"`
	Children []Block    `json:"children,omitempty"`
	Icon     *Icon      `json:"icon,omitempty"`
	Color    Color      `json:"color,omitempty"`
}

type CodeBlock struct {
	RichText []RichText `json:"rich_text"`
	Children []Block    `json:"children,omitempty"`
	Caption  []RichText `json:"caption,omitempty"`
	Language *string    `json:"language,omitempty"`
}

type EmbedBlock struct {
	URL string `json:"url"`
}

type ImageBlock struct {
	Type     FileType      `json:"type"`
	File     *FileFile     `json:"file,omitempty"`
	External *FileExternal `json:"external,omitempty"`
	Caption  []RichText    `json:"caption,omitempty"`
}

type AudioBlock struct {
	Type     FileType      `json:"type"`
	File     *FileFile     `json:"file,omitempty"`
	External *FileExternal `json:"external,omitempty"`
	Caption  []RichText    `json:"caption,omitempty"`
}

type VideoBlock struct {
	Type     FileType      `json:"type"`
	File     *FileFile     `json:"file,omitempty"`
	External *FileExternal `json:"external,omitempty"`
	Caption  []RichText    `json:"caption,omitempty"`
}

type FileBlock struct {
	Type     FileType      `json:"type"`
	File     *FileFile     `json:"file,omitempty"`
	External *FileExternal `json:"external,omitempty"`
	Caption  []RichText    `json:"caption,omitempty"`
}

type PDFBlock struct {
	Type     FileType      `json:"type"`
	File     *FileFile     `json:"file,omitempty"`
	External *FileExternal `json:"external,omitempty"`
	Caption  []RichText    `json:"caption,omitempty"`
}

type BookmarkBlock struct {
	URL     string     `json:"url"`
	Caption []RichText `json:"caption,omitempty"`
}

type EquationBlock struct {
	Expression string `json:"expression"`
}

type ColumnListBlock struct {
	Children []ColumnBlock `json:"children,omitempty"`
}

type ColumnBlock struct {
	Children []Block `json:"children,omitempty"`
}

type TableBlock struct {
	TableWidth      int     `json:"table_width"`
	HasColumnHeader bool    `json:"has_column_header"`
	HasRowHeader    bool    `json:"has_row_header"`
	Children        []Block `json:"children,omitempty"`
}

type TableRowBlock struct {
	Cells [][]RichText `json:"cells"`
}

type LinkPreviewBlock struct {
	URL string `json:"url"`
}

type LinkToPageBlock struct {
	Type       LinkToPageType `json:"type"`
	PageID     string         `json:"page_id,omitempty"`
	DatabaseID string         `json:"database_id,omitempty"`
}

type LinkToPageType string

const (
	LinkToPageTypePageID     LinkToPageType = "page_id"
	LinkToPageTypeDatabaseID LinkToPageType = "database_id"
)

type SyncedBlock struct {
	BaseBlock

	SyncedFrom *SyncedFrom `json:"synced_from"`
	Children   []Block     `json:"children,omitempty"`
}

// MarshalJSON implements json.Marshaler.
func (b SyncedBlock) MarshalJSON() ([]byte, error) {
	type (
		blockAlias SyncedBlock
		dto        struct {
			SyncedBlock blockAlias `json:"synced_block"`
		}
	)

	return json.Marshal(dto{
		SyncedBlock: blockAlias(b),
	})
}

type SyncedFrom struct {
	Type    SyncedFromType `json:"type"`
	BlockID string         `json:"block_id"`
}

type SyncedFromType string

const SyncedFromTypeBlockID SyncedFromType = "block_id"

type DividerBlock struct {
	BaseBlock
}

// MarshalJSON implements json.Marshaler.
func (b DividerBlock) MarshalJSON() ([]byte, error) {
	type (
		blockAlias DividerBlock
		dto        struct {
			Divider blockAlias `json:"divider"`
		}
	)

	return json.Marshal(dto{
		Divider: blockAlias(b),
	})
}

type TableOfContentsBlock struct {
	BaseBlock

	Color Color `json:"color,omitempty"`
}

// MarshalJSON implements json.Marshaler.
func (b TableOfContentsBlock) MarshalJSON() ([]byte, error) {
	type (
		blockAlias TableOfContentsBlock
		dto        struct {
			TableOfContents blockAlias `json:"table_of_contents"`
		}
	)

	return json.Marshal(dto{
		TableOfContents: blockAlias(b),
	})
}

type BreadcrumbBlock struct {
	BaseBlock
}

// MarshalJSON implements json.Marshaler.
func (b BreadcrumbBlock) MarshalJSON() ([]byte, error) {
	type (
		blockAlias BreadcrumbBlock
		dto        struct {
			Breadcrumb blockAlias `json:"breadcrumb"`
		}
	)

	return json.Marshal(dto{
		Breadcrumb: blockAlias(b),
	})
}

type BlockType string

const (
	BlockTypeParagraph        BlockType = "paragraph"
	BlockTypeHeading1         BlockType = "heading_1"
	BlockTypeHeading2         BlockType = "heading_2"
	BlockTypeHeading3         BlockType = "heading_3"
	BlockTypeBulletedListItem BlockType = "bulleted_list_item"
	BlockTypeNumberedListItem BlockType = "numbered_list_item"
	BlockTypeToDo             BlockType = "to_do"
	BlockTypeToggle           BlockType = "toggle"
	BlockTypeChildPage        BlockType = "child_page"
	BlockTypeChildDatabase    BlockType = "child_database"
	BlockTypeCallout          BlockType = "callout"
	BlockTypeQuote            BlockType = "quote"
	BlockTypeCode             BlockType = "code"
	BlockTypeEmbed            BlockType = "embed"
	BlockTypeImage            BlockType = "image"
	BlockTypeAudio            BlockType = "audio"
	BlockTypeVideo            BlockType = "video"
	BlockTypeFile             BlockType = "file"
	BlockTypePDF              BlockType = "pdf"
	BlockTypeBookmark         BlockType = "bookmark"
	BlockTypeEquation         BlockType = "equation"
	BlockTypeDivider          BlockType = "divider"
	BlockTypeTableOfContents  BlockType = "table_of_contents"
	BlockTypeBreadCrumb       BlockType = "breadcrumb"
	BlockTypeColumnList       BlockType = "column_list"
	BlockTypeColumn           BlockType = "column"
	BlockTypeTable            BlockType = "table"
	BlockTypeTableRow         BlockType = "table_row"
	BlockTypeLinkPreview      BlockType = "link_preview"
	BlockTypeLinkToPage       BlockType = "link_to_page"
	BlockTypeSyncedBlock      BlockType = "synced_block"
	BlockTypeTemplate         BlockType = "template"
	BlockTypeUnsupported      BlockType = "unsupported"
)

type PaginationQuery struct {
	StartCursor string
	PageSize    int
}

// BlockChildrenResponse contains results (block children) and pagination data returned from a find request.
type BlockChildrenResponse struct {
	Results    []Block
	HasMore    bool
	NextCursor *string
}

func (resp *BlockChildrenResponse) UnmarshalJSON(b []byte) error {
	type responseDTO struct {
		Results    []BlockDTO `json:"results"`
		HasMore    bool       `json:"has_more"`
		NextCursor *string    `json:"next_cursor"`
	}

	var dto responseDTO

	if err := json.Unmarshal(b, &dto); err != nil {
		return err
	}

	resp.HasMore = dto.HasMore
	resp.NextCursor = dto.NextCursor
	resp.Results = make([]Block, len(dto.Results))

	for i, blockDTO := range dto.Results {
		resp.Results[i] = blockDTO
	}

	return nil
}

func (dto BlockDTO) Block() Block {
	return dto
}
