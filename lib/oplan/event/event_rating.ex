defmodule Oplan.Event.EventRating do
  use Ecto.Schema
  import Ecto.Changeset

  alias Oplan.Account.User
  alias Oplan.Event.Event

  schema "event_rating" do
    field :rating, :integer
    belongs_to :user, User
    belongs_to :event, Event

    timestamps()
  end

  @doc false
  def changeset(event_rating, attrs) do
    event_rating
    |> cast(attrs, [:user_id, :event_id, :rating])
    |> validate_required([:user_id, :event_id, :rating])
  end
end
