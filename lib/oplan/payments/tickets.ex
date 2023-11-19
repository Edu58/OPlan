defmodule Oplan.Payments.Tickets do
  use Ecto.Schema
  import Ecto.Changeset

  alias Oplan.Account.User
  alias Oplan.Event.Event

  schema "event_tickets" do
    field :accepted_mode_of_payment, {:array, :string}
    field :currency, :string
    field :price, :integer
    field :valid_from, :utc_datetime
    field :expiry, :utc_datetime
    field :used, :boolean, default: false
    belongs_to :user, User
    belongs_to :event, Event

    timestamps()
  end

  @doc false
  def changeset(tickets, attrs) do
    tickets
    |> cast(attrs, [:user_id, :event_id, :price, :currency, :accepted_mode_of_payment])
    |> validate_required([:user_id, :event_id, :price, :currency, :accepted_mode_of_payment])
  end
end
