defmodule Oplan.Event.Event do
  use Ecto.Schema
  import Ecto.Changeset

  alias Oplan.Account.User
  alias Oplan.Event.EventType
  alias Oplan.Event.EventGuest
  alias Oplan.Event.EventRating
  alias Oplan.Event.EventComment
  alias Oplan.Payments.Tickets

  schema "events" do
    field :age_restriction, :integer
    field :banner, :string
    field :description, :string
    field :end, :utc_datetime
    field :free, :boolean, default: false
    field :name, :string
    field :notify_attendess, :boolean, default: false
    field :photos, :string
    field :policies_and_rules, :string
    field :public, :boolean, default: false
    field :start, :utc_datetime
    field :ticket_price, :string
    field :venue, :string
    field :number_of_tickets, :integer
    field :sponsors, {:array, :map}
    belongs_to :user, User
    belongs_to :event_type, EventType

    has_many :event_ratings, EventRating, on_replace: :delete

    has_many :event_comments, EventComment, on_replace: :delete

    has_many :event_tickets, Tickets, on_replace: :delete

    many_to_many :guests, EventGuest, join_through: "event_guests"

    timestamps()
  end

  @doc false
  def changeset(event, attrs) do
    event
    |> cast(attrs, [
      :name,
      :venue,
      :start,
      :end,
      :description,
      :public,
      :free,
      :ticket_price,
      :banner,
      :photos,
      :age_restriction,
      :policies_and_rules,
      :notify_attendess,
      :sponsors,
      :number_of_tickets
    ])
    |> validate_required([
      :name,
      :venue,
      :start,
      :end,
      :description
    ])
  end
end
