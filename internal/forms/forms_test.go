package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestFormValid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestFormRequired(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.Form)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("should not have required fields when it does")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("should have required fields when it does")
	}
}

func TestFormHas(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("form shows has field when it does not should")
	}
}

func TestMinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("for shows min length for non-existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedValues := url.Values{}
	postedValues.Add("some_field", "some_value")
	form = New(postedValues)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows minLength of 100 met when data is shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "abc123")
	form = New(postedValues)
	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows minLength of 1 is not met if it is")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have an error, but did get one")
	}
}

func TestFormIsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("for shows email for non-existent field")
	}

	postedValues := url.Values{}
	postedValues.Add("email", "abc123@hear.com")
	form = New(postedValues)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("shows email field failed")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "x")
	form = New(postedValues)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("shows email x not email failed")
	}
}
