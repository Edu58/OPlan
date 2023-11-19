defmodule Oplan.Repo.Migrations.CreateEvent do
  use Ecto.Migration

  def change do
    create table(:events) do
      add :name, :string
      add :venue, :string
      add :start, :utc_datetime
      add :end, :utc_datetime
      add :description, :text
      add :public, :boolean, default: false, null: false
      add :free, :boolean, default: false, null: false
      add :ticket_price, :string
      add :banner, :string
      add :photos, :string
      add :age_restriction, :integer
      add :policies_and_rules, :text
      add :notify_attendess, :boolean, default: false, null: false
      add :sponsors, {:array, :map}
      add :number_of_tickets, :integer
      add :user_id, references(:users, on_delete: :delete_all), null: false
      add :event_type_id, references(:event_types, on_delete: :delete_all), null: false

      timestamps()
    end

    create index(:events, [:name, :venue])
    create unique_index(:events, [:event_type_id])
    create unique_index(:events, [:name, :user_id])
  end
end
