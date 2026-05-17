package utils

import (
	"fmt"
	"simpleorder/internal/domain"

	"github.com/go-pdf/fpdf"
)

func GenerateInvoicePDF(order *domain.Order) (*fpdf.Fpdf, error) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	// Header
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(190, 10, "INVOICE", "0", 1, "C", false, 0, "")
	pdf.Ln(10)

	// Company info and Invoice Details
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(95, 8, "SimpleOrder Inc.", "0", 0, "L", false, 0, "")
	pdf.CellFormat(95, 8, fmt.Sprintf("Invoice #: %s", order.InvoiceNumber), "0", 1, "R", false, 0, "")
	
	pdf.CellFormat(95, 8, "123 Business Road", "0", 0, "L", false, 0, "")
	pdf.CellFormat(95, 8, fmt.Sprintf("Date: %s", order.CreatedAt.Format("2006-01-02")), "0", 1, "R", false, 0, "")
	pdf.Ln(10)

	// Customer Info
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(190, 8, "Bill To:", "0", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(190, 8, fmt.Sprintf("Name: %s", order.Customer.Name), "0", 1, "L", false, 0, "")
	pdf.CellFormat(190, 8, fmt.Sprintf("Email: %s", order.Customer.Email), "0", 1, "L", false, 0, "")
	if order.Customer.Phone != "" {
		pdf.CellFormat(190, 8, fmt.Sprintf("Phone: %s", order.Customer.Phone), "0", 1, "L", false, 0, "")
	}
	if order.Customer.Address != "" {
		pdf.CellFormat(190, 8, fmt.Sprintf("Address: %s", order.Customer.Address), "0", 1, "L", false, 0, "")
	}
	pdf.Ln(10)

	// Items Table Header
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(80, 10, "Description", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 10, "Quantity", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Unit Price", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Subtotal", "1", 1, "C", false, 0, "")

	// Items Table Body
	pdf.SetFont("Arial", "", 12)
	for _, item := range order.Items {
		pdf.CellFormat(80, 10, item.Product.Name, "1", 0, "L", false, 0, "")
		pdf.CellFormat(30, 10, fmt.Sprintf("%d", item.Quantity), "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("$%.2f", item.Price), "1", 0, "R", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("$%.2f", item.SubTotal), "1", 1, "R", false, 0, "")
	}

	// Total
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(150, 10, "Total", "1", 0, "R", false, 0, "")
	pdf.CellFormat(40, 10, fmt.Sprintf("$%.2f", order.TotalAmount), "1", 1, "R", false, 0, "")

	return pdf, nil
}
