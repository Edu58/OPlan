defmodule Oplan.SupportTest do
  use Oplan.DataCase

  alias Oplan.Support

  describe "support_tickets" do
    alias Oplan.Support.SupportTickets

    import Oplan.SupportFixtures

    @invalid_attrs %{id: nil, message: nil}

    test "list_support_tickets/0 returns all support_tickets" do
      support_tickets = support_tickets_fixture()
      assert Support.list_support_tickets() == [support_tickets]
    end

    test "get_support_tickets!/1 returns the support_tickets with given id" do
      support_tickets = support_tickets_fixture()
      assert Support.get_support_tickets!(support_tickets.id) == support_tickets
    end

    test "create_support_tickets/1 with valid data creates a support_tickets" do
      valid_attrs = %{id: "7488a646-e31f-11e4-aace-600308960662", message: "some message"}

      assert {:ok, %SupportTickets{} = support_tickets} = Support.create_support_tickets(valid_attrs)
      assert support_tickets.id == "7488a646-e31f-11e4-aace-600308960662"
      assert support_tickets.message == "some message"
    end

    test "create_support_tickets/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Support.create_support_tickets(@invalid_attrs)
    end

    test "update_support_tickets/2 with valid data updates the support_tickets" do
      support_tickets = support_tickets_fixture()
      update_attrs = %{id: "7488a646-e31f-11e4-aace-600308960668", message: "some updated message"}

      assert {:ok, %SupportTickets{} = support_tickets} = Support.update_support_tickets(support_tickets, update_attrs)
      assert support_tickets.id == "7488a646-e31f-11e4-aace-600308960668"
      assert support_tickets.message == "some updated message"
    end

    test "update_support_tickets/2 with invalid data returns error changeset" do
      support_tickets = support_tickets_fixture()
      assert {:error, %Ecto.Changeset{}} = Support.update_support_tickets(support_tickets, @invalid_attrs)
      assert support_tickets == Support.get_support_tickets!(support_tickets.id)
    end

    test "delete_support_tickets/1 deletes the support_tickets" do
      support_tickets = support_tickets_fixture()
      assert {:ok, %SupportTickets{}} = Support.delete_support_tickets(support_tickets)
      assert_raise Ecto.NoResultsError, fn -> Support.get_support_tickets!(support_tickets.id) end
    end

    test "change_support_tickets/1 returns a support_tickets changeset" do
      support_tickets = support_tickets_fixture()
      assert %Ecto.Changeset{} = Support.change_support_tickets(support_tickets)
    end
  end
end
