defmodule Oplan.PaymentsTest do
  use Oplan.DataCase

  alias Oplan.Payments

  describe "event_tickets" do
    alias Oplan.Payments.Tickets

    import Oplan.PaymentsFixtures

    @invalid_attrs %{accepted_mode_of_payment: nil, currency: nil, price: nil}

    test "list_event_tickets/0 returns all event_tickets" do
      tickets = tickets_fixture()
      assert Payments.list_event_tickets() == [tickets]
    end

    test "get_tickets!/1 returns the tickets with given id" do
      tickets = tickets_fixture()
      assert Payments.get_tickets!(tickets.id) == tickets
    end

    test "create_tickets/1 with valid data creates a tickets" do
      valid_attrs = %{
        accepted_mode_of_payment: ["option1", "option2"],
        currency: "some currency",
        price: 42
      }

      assert {:ok, %Tickets{} = tickets} = Payments.create_tickets(valid_attrs)
      assert tickets.accepted_mode_of_payment == ["option1", "option2"]
      assert tickets.currency == "some currency"
      assert tickets.price == 42
    end

    test "create_tickets/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Payments.create_tickets(@invalid_attrs)
    end

    test "update_tickets/2 with valid data updates the tickets" do
      tickets = tickets_fixture()

      update_attrs = %{
        accepted_mode_of_payment: ["option1"],
        currency: "some updated currency",
        price: 43
      }

      assert {:ok, %Tickets{} = tickets} = Payments.update_tickets(tickets, update_attrs)
      assert tickets.accepted_mode_of_payment == ["option1"]
      assert tickets.currency == "some updated currency"
      assert tickets.price == 43
    end

    test "update_tickets/2 with invalid data returns error changeset" do
      tickets = tickets_fixture()
      assert {:error, %Ecto.Changeset{}} = Payments.update_tickets(tickets, @invalid_attrs)
      assert tickets == Payments.get_tickets!(tickets.id)
    end

    test "delete_tickets/1 deletes the tickets" do
      tickets = tickets_fixture()
      assert {:ok, %Tickets{}} = Payments.delete_tickets(tickets)
      assert_raise Ecto.NoResultsError, fn -> Payments.get_tickets!(tickets.id) end
    end

    test "change_tickets/1 returns a tickets changeset" do
      tickets = tickets_fixture()
      assert %Ecto.Changeset{} = Payments.change_tickets(tickets)
    end
  end
end
