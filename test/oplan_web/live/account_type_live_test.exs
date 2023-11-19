defmodule OplanWeb.AccountTypeLiveTest do
  use OplanWeb.ConnCase

  import Phoenix.LiveViewTest
  import Oplan.AccountFixtures

  @create_attrs %{name: "some name"}
  @update_attrs %{name: "some updated name"}
  @invalid_attrs %{name: nil}

  defp create_account_type(_) do
    account_type = account_type_fixture()
    %{account_type: account_type}
  end

  describe "Index" do
    setup [:create_account_type]

    test "lists all account_types", %{conn: conn, account_type: account_type} do
      {:ok, _index_live, html} = live(conn, ~p"/account_types")

      assert html =~ "Listing Account types"
      assert html =~ account_type.name
    end

    test "saves new account_type", %{conn: conn} do
      {:ok, index_live, _html} = live(conn, ~p"/account_types")

      assert index_live |> element("a", "New Account type") |> render_click() =~
               "New Account type"

      assert_patch(index_live, ~p"/account_types/new")

      assert index_live
             |> form("#account_type-form", account_type: @invalid_attrs)
             |> render_change() =~ "can&#39;t be blank"

      assert index_live
             |> form("#account_type-form", account_type: @create_attrs)
             |> render_submit()

      assert_patch(index_live, ~p"/account_types")

      html = render(index_live)
      assert html =~ "Account type created successfully"
      assert html =~ "some name"
    end

    test "updates account_type in listing", %{conn: conn, account_type: account_type} do
      {:ok, index_live, _html} = live(conn, ~p"/account_types")

      assert index_live
             |> element("#account_types-#{account_type.id} a", "Edit")
             |> render_click() =~
               "Edit Account type"

      assert_patch(index_live, ~p"/account_types/#{account_type}/edit")

      assert index_live
             |> form("#account_type-form", account_type: @invalid_attrs)
             |> render_change() =~ "can&#39;t be blank"

      assert index_live
             |> form("#account_type-form", account_type: @update_attrs)
             |> render_submit()

      assert_patch(index_live, ~p"/account_types")

      html = render(index_live)
      assert html =~ "Account type updated successfully"
      assert html =~ "some updated name"
    end

    test "deletes account_type in listing", %{conn: conn, account_type: account_type} do
      {:ok, index_live, _html} = live(conn, ~p"/account_types")

      assert index_live
             |> element("#account_types-#{account_type.id} a", "Delete")
             |> render_click()

      refute has_element?(index_live, "#account_types-#{account_type.id}")
    end
  end

  describe "Show" do
    setup [:create_account_type]

    test "displays account_type", %{conn: conn, account_type: account_type} do
      {:ok, _show_live, html} = live(conn, ~p"/account_types/#{account_type}")

      assert html =~ "Show Account type"
      assert html =~ account_type.name
    end

    test "updates account_type within modal", %{conn: conn, account_type: account_type} do
      {:ok, show_live, _html} = live(conn, ~p"/account_types/#{account_type}")

      assert show_live |> element("a", "Edit") |> render_click() =~
               "Edit Account type"

      assert_patch(show_live, ~p"/account_types/#{account_type}/show/edit")

      assert show_live
             |> form("#account_type-form", account_type: @invalid_attrs)
             |> render_change() =~ "can&#39;t be blank"

      assert show_live
             |> form("#account_type-form", account_type: @update_attrs)
             |> render_submit()

      assert_patch(show_live, ~p"/account_types/#{account_type}")

      html = render(show_live)
      assert html =~ "Account type updated successfully"
      assert html =~ "some updated name"
    end
  end
end
